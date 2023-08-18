package controllers

import (
	"net/http"

	ai "meg-server/ai"
	models "meg-server/models"
	"meg-server/utils"

	"context"

	"github.com/gin-gonic/gin"

	openai "github.com/sashabaranov/go-openai"

	_ "github.com/joho/godotenv/autoload"
)


func Chat(c *gin.Context) {
	// Get the OpenAI client
	client := ai.GetClient()
	ctx := context.Background()
	b, _ := utils.ReadRequestBody(c)
	Messages :=[]openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: string(b)}}
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: Messages,
		MaxTokens: 10,
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	// Error from OpenAI ChatComletion API
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error() ,
		})

		return
	}
	// Format chat response object base on model
	var ChatResponse = models.ChatResponse{
		Message: resp.Choices[0].Message.Content,
	}
	// Respond with the message from OpenAI
	c.JSON(http.StatusOK, ChatResponse)


}