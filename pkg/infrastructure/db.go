package infrastructure

import (
	"Restro/pkg/framework"
	"fmt"

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
	if env.DBType != "mysql" {
		url = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/?charset=utf8mb4&parseTime=True&loc=Local",
			env.DBUsername,
			env.DBPassword,
			env.DBHost,
		)
	}

	logger.Info("opening db connection")
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: logger.GetGormLogger()})
	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	logger.Info("creating database if it does't exist")
	if err = db.Exec("CREATE DATABASE IF NOT EXISTS " + env.DBName).Error; err != nil {
		logger.Info("couldn't create database")
		logger.Panic(err)
	}

	logger.Info("using given database")
	if err := db.Exec(fmt.Sprintf("USE %s", env.DBName)).Error; err != nil {
		logger.Info("cannot use the given database")
		logger.Panic(err)
	}
	logger.Info("database connection established")

	logger.Info("currentDatabase:", db.Migrator().CurrentDatabase())


	return Database{DB: db}
}