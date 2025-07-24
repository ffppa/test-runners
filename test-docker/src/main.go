// main.go
package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response structure for mock response data
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Token   string `json:"token"` // Add token field to the response
}

func main() {
	// Print the startup banner
	printBanner()

	http.HandleFunc("/mock", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Generate a random token
		token, err := generateRandomToken(16) // Generates a 16-byte (32-char) token
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Prepare a mock response
		response := Response{
			Message: "Example of a mock API!",
			Status:  "success",
			Token:   token,
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

// Function to print a startup banner
func printBanner() {
	fmt.Println(`
==========================================
         MOCK API SERVER STARTED
==========================================
         Access the API at:
      -> http://localhost:8080/mock
==========================================
	`)
}

// Function to generate a random token
func generateRandomToken(length int) (string, error) {
	bytes := make([]byte, length) // Create a byte slice of the desired length
	_, err := rand.Read(bytes)   // Fill the slice with random bytes
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a hexadecimal string
	return hex.EncodeToString(bytes), nil
}