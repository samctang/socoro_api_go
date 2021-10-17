package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samctang/socoro_api/controllers"
	"github.com/samctang/socoro_api/database"
	"os"
)

func main() {
	os.Setenv("PORT", "8080")

	database.Client()
	defer database.Conn.Close()

	r := gin.Default()

	r.GET("/operation", controllers.GetOperations)

	r.Run()
}
