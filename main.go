package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samctang/socoro_api/database"
	"github.com/samctang/socoro_api/operation"
	"github.com/samctang/socoro_api/user"
	"os"
)

func main() {
	os.Setenv("PORT", "8080")

	// Set up DB and terminate connection at the end
	database.SetupDB()
	defer database.Conn.Close()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.CORS()) //remove this for prod deployment
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/user/register", user.Register)
	e.POST("/user/login", user.Login)
	e.GET("/operation", operation.GetOperations)

	e.Logger.Fatal(e.Start(":1323"))
}
