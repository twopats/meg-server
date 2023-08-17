package controllers

import (
	"net/http"

	ai "meg-server/ai"

	"context"

	"github.com/gin-gonic/gin"

	openai "github.com/sashabaranov/go-openai"

	_ "github.com/joho/godotenv/autoload"
)


func Chat(c *gin.Context) {
	client := ai.GetClient()
	ctx := context.Background()

	reqContent := c.Query("content")

	Messages :=[]openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: reqContent}}


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
	

	// Respond with the message from OpenAI
	c.JSON(http.StatusOK, gin.H{
		"message": resp.Choices[0].Message.Content,
	})


}