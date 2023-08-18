package router

import (
	"meg-server/controllers"

	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
  r := gin.Default()
 
  r.GET("/", controllers.Home)
  r.GET("/api/ping", controllers.Ping)
  r.POST("/api/chat", controllers.Chat)
  r.GET("/api/ws-chat", controllers.WS_Chat)
  r.POST("/api/image", controllers.Image)
  

  return r;

}