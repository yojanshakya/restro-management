package main

import (
    "Restro/bootstrap"

    "github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()
    _ = bootstrap.RootApp.Execute()
}
