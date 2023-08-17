package tests

import (
	"bytes"
	"encoding/json"
	"meg-server/models"
	router "meg-server/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)




func TestPingRoute(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/ping", nil)
	router.ServeHTTP(w, req)

	mockResponse := models.ChatResponse{Message: "pong"}

	b,_ := json.Marshal(mockResponse)
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(b), w.Body.String())
}

// TODO: Fix this test - can we test the API call to OpenAI via the key?
func TestChatRoute(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	mockRequest := models.ChatRequest{Content: "How many colors on a rainbow?"}
	jsonValue, _ := json.Marshal(mockRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/chat",bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	
	
	t.Log(req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	// assert.Equal(t, string(b), w.Body.String())
}
