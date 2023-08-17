package main

import (
	"log"
	router "meg-server/router"
	"os"

	"github.com/gorilla/websocket"
	_ "github.com/joho/godotenv/autoload"
)




var upgrader = websocket.Upgrader{
 ReadBufferSize:  1024,
 WriteBufferSize: 1024,
}


func main() {
  r := router.SetupRouter()

  port := os.Getenv("PORT")
    if port == "" {
    log.Fatal("PORT not set in .env file")
  }

  
  r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}



