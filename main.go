package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
  "meg-server/controllers"
)

func main() {
  // router
  r := gin.Default()

  r.GET("/api/ping", controllers.Ping)

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

