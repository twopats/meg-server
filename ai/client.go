package ai

import (
	"os"

	openai "github.com/sashabaranov/go-openai"
)


var c *openai.Client

func init() {
	c = openai.NewClient(os.Getenv("OPENAI_API_KEY"))

}

func GetClient() *openai.Client {
	return c
}