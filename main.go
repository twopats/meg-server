package main

import (
	"log"
	"meg-server/controllers"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
  // router
  r := gin.Default()


  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("PORT not set in .env file")
  }

  r.GET("/api/ping", controllers.Ping)

  r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

