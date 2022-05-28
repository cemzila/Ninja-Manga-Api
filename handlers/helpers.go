package handlers

import (
	"encoding/json"
	"net/http"
)

// Function for Sending Response
func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// type CreateMangaRequestPayload struct {
// 	Title       string `json:"title"`
// 	Cover       string `json:"cover"`
// 	Artist      string `json:"artist"`
// 	Description string `json:"description"`
// 	Tags        string `json:"tags"`
// }
