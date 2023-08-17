package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
 ReadBufferSize:  1024,
 WriteBufferSize: 1024,
}

func WS_Chat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	
		if err != nil {
			println(err.Error())
			return
		}
		defer conn.Close()		
                i := 0
		for {
			i++
			conn.WriteMessage(websocket.TextMessage, []byte("New Message (#"+strconv.Itoa(i)+")"))
			time.Sleep(2 * time.Second)
		}

}