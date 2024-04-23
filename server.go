package main

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/public/index.html")
}

func GenerateTextHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for generating text from the given song
	// Call AI model, preprocess data, handle response, etc.
	// Example:
	// generatedText := generateTextFromAI(songData)

	// Dummy response for testing
	generatedText := "Generated lyrics will be here..."

	// Return generated text as response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(generatedText))
}
