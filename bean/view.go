package bean

import (
	"encoding/json"
	"net/http"
	"time"
)

type APIResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Signature interface{} `json:"signature"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}

func JsonResponse(w http.ResponseWriter, statusCode int, status string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := APIResponse{
		Status:    status,
		Message:   message,
		Signature: nil,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339Nano),
	}

	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	JsonResponse(w, statusCode, "99", message, nil)
}
