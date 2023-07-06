package main

import (
	"biodata/database"
	"biodata/routes"
	"biodata/utils"

	"github.com/labstack/echo"
)

func main() {
	config := database.Config{
		DB_USERNAME: utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: utils.GetConfig("DB_PASSWORD"),
		DB_HOST:     utils.GetConfig("DB_HOST"),
		DB_PORT:     utils.GetConfig("DB_PORT"),
		DB_NAME:     utils.GetConfig("DB_NAME"),
	}

	config.ConnectDB()

	database.MigrateDB()

	e := echo.New()
	routes.SetUpRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
