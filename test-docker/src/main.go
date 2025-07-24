// main.go
package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

  http.Handle("/mock", http.StripPrefix("/mock", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

    // Log the request details in a fancy style
		logRequest(r, token)

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
====================================================
███╗   ███╗ ██████╗  ██████╗██╗  ██╗     █████╗ ██████╗ 
████╗ ████║██╔═══██╗██╔════╝██║  ██║    ██╔══██╗██╔══██╗
██╔████╔██║██║   ██║██║     ███████║    ███████║██████╔╝
██║╚██╔╝██║██║   ██║██║     ██╔══██║    ██╔══██║██╔═══╝ 
██║ ╚═╝ ██║╚██████╔╝╚██████╗██║  ██║    ██║  ██║██║     
╚═╝     ╚═╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝    ╚═╝  ╚═╝╚═╝     
====================================================
         MOCK API SERVER - READY TO SERVE
         Access the API at:
      -> http://localhost:8080/mock
====================================================
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

// Function to log incoming requests in a "mainframe" style
func logRequest(r *http.Request, token string) {
	clientIP := r.RemoteAddr // Get the client IP address
	method := r.Method       // Get the HTTP method
	url := r.URL.String()    // Get the requested URL
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Fancy log output
	fmt.Printf(`
===================================================
🚀 NEW REQUEST RECEIVED
====================================================
📅 TIMESTAMP: %s
🌐 CLIENT IP: %s
🔍 REQUESTED URL: %s
🔧 METHOD: %s
🔐 TOKEN: %s
====================================================
`, timestamp, clientIP, url, method, token)
}