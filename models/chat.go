package models

type ChatResponse struct {
	Message string `json:"message"`
}

type ChatRequest struct {
	Content string `json:"content"`
}