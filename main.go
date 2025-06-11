package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// pongResponse represents the JSON structure for /ping
type pongResponse struct {
	Message string `json:"message"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	resp := pongResponse{Message: "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var payload map[string]interface{}
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, `{"error": "invalid JSON"}`, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(payload)
}

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			pingHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			echoHandler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}