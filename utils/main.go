package utils

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ReadRequestBody(c *gin.Context) ([]byte, error) {
	b,bErr := io.ReadAll(c.Request.Body)
	if bErr != nil {
		c.Status(http.StatusBadRequest)
		return nil, bErr
	}

	return b, nil
}