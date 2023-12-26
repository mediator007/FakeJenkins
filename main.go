package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		// Extract id from the URL path
		id := extractID(r.URL.Path)

		// Check if id is empty
		if id == "" {
			http.Error(w, "Invalid URL pattern. Missing ID.", http.StatusBadRequest)
			return
		}

		// Respond with the extracted ID in JSON format
		response := fmt.Sprintf(`{"id": "%s"}`, id)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})

	// Start the server on port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func extractID(path string) string {
	// Split the path using "/"
	parts := strings.Split(path, "/")

	// Check if there are enough parts and the second part is "api"
	if len(parts) >= 3 && parts[1] == "api" {
		return parts[2] // The third part is the ID
	}

	return ""
}
