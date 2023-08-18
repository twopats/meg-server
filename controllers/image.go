package controllers

import (
	"meg-server/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Image(c *gin.Context) {
	// b,_ := utils.ReadRequestBody(c)
	// http.NewRequest(http.MethodPost, config.GetConfig().External.Image, b)
	_, err := http.Post(config.GetConfig().External.Image, "application/json", nil)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	//TODO: format image response
	c.JSON(http.StatusOK, "TODO: implement image generation pipeline")
}
