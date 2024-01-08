package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		http.ServeFile(w, r, "index.html")
		return
	}

	if r.Method != http.MethodPost {
		sendJsonResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var requestData map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		sendJsonResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if message, ok := requestData["message"].(string); !ok || message == "" {
		sendJsonResponse(w, http.StatusBadRequest, "Invalid JSON message")
		return
	}

	fmt.Println("Received message:", requestData["message"])

	response := JsonResponse{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func sendJsonResponse(w http.ResponseWriter, statusCode int, message string) {
	response := JsonResponse{
		Status:  fmt.Sprintf("%d", statusCode),
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
