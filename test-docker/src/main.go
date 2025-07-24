// main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response structure for mock response data
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Prepare a mock response
		response := Response{
			Message: "ok",
			Status:  "success",
		}

		// Encode the response to JSON and write it to the ResponseWriter
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/mock", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Prepare a mock response
		response := Response{
			Message: "Example of a Mock API!",
			Status:  "success",
		}

		// Encode the response to JSON and write it to the ResponseWriter
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		}
	})

	// Start the server on port 8080
	port := "8080"
	log.Printf("Server is running at http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}