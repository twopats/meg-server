package main

import (
	"log"
	controllers "meg-server/controllers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "github.com/joho/godotenv/autoload"
)

var upgrader = websocket.Upgrader{
 ReadBufferSize:  1024,
 WriteBufferSize: 1024,
}

func main() {



  r := gin.Default()



  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("PORT not set in .env file")
  }

  r.GET("/api/ping", controllers.Ping)
  r.GET("/api/chat", controllers.Chat)

  r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

