package main

import (
	config "meg-server/config"
	router "meg-server/router"

	"github.com/gorilla/websocket"
	_ "github.com/joho/godotenv/autoload"
)




var upgrader = websocket.Upgrader{
 ReadBufferSize:  1024,
 WriteBufferSize: 1024,
}


func main() {
  r := router.SetupRouter()


  
  r.Run(":" + config.GetConfig().Server.Port) 
}



