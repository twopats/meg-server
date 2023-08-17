package router

import (
	"meg-server/controllers"

	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
  r := gin.Default()
 
  r.GET("/api/ping", controllers.Ping)
  r.POST("/api/chat", controllers.Chat)

  return r;

}