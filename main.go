package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samctang/socoro_api/controllers"
	"github.com/samctang/socoro_api/database"
	"os"
	"time"
)

func main() {
	os.Setenv("PORT", "8080")

	// Start DB Client and terminate connection at the end
	database.Client()
	defer database.Conn.Close()

	r := gin.Default()

	// CORS for localhost and https://github.com origins, allowing:
	// - * methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/operation", controllers.GetOperations)

	r.Run()
}
