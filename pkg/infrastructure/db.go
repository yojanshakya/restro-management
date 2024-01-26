package infrastructure

import (
	"Restro/pkg/framework"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(logger framework.Logger, env *framework.Env) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort)

	logger.Info("opening db connection")
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: logger.GetGormLogger()})
	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	logger.Info("creating database if it doesn't exist")
	if err = db.Exec("CREATE DATABASE IF NOT EXISTS `" + env.DBName + "`").Error; err != nil {
		logger.Info("couldn't create database")
		logger.Panic(err)
	}

	// close the current connection
	sqlDb, err := db.DB()
	if err != nil {
		logger.Panic(err)
	}
	if dbErr := sqlDb.Close(); dbErr != nil {
		logger.Panic(err)
	}

	// reopen connection with the given database, after creating or checking if the database exists
	logger.Info("using given database")
	urlWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	db, err = gorm.Open(mysql.Open(urlWithDB), &gorm.Config{Logger: logger.GetGormLogger()})
	if err != nil {
		logger.Info("Url: ", urlWithDB)
		logger.Panic(err)
	}

	conn, err := db.DB()
	if err != nil {
		logger.Info("couldn't get db connection")
		logger.Panic(err)
	}

	conn.SetConnMaxLifetime(time.Minute * 5)
	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(1)

	return Database{DB: db}
}