package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// EmailData represents the dynamic data for the email template
type EmailData struct {
	LogoURL       string `json:"logo_url"`
	Title         string `json:"title"`
	Users         []User `json:"users"`
	ReceiverEmail string `json:"receiver_email"`
	LegalAddress  string `json:"legal_address"`
	BodyColor     string `json:"bodyColor"`
	PrimaryColor  string `json:"primaryColor"`
	ButtonColor   string `json:"buttonColor"`
	SenderName    string `json:"sender_name"`
}

// User represents user details
type User struct {
	UserName string `json:"user_name"`
	UserCity string `json:"user_city"`
	PhotoURL string `json:"photo_url"`
}

func main() {
	// Set up the HTTP server to serve the template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Parse the email template
		tmplPath := filepath.Join("templates", "test-email-template.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Unable to parse template", http.StatusInternalServerError)
			log.Fatalf("Error parsing template: %v", err)
			return
		}

		// Load the JSON data for the template
		dataFile := filepath.Join("data", "emailData.json")
		file, err := os.Open(dataFile)
		if err != nil {
			http.Error(w, "Unable to open data file", http.StatusInternalServerError)
			log.Fatalf("Error opening JSON file: %v", err)
			return
		}
		defer file.Close()

		// Decode JSON data into the EmailData struct
		var emailData EmailData
		if err := json.NewDecoder(file).Decode(&emailData); err != nil {
			http.Error(w, "Unable to decode JSON", http.StatusInternalServerError)
			log.Fatalf("Error decoding JSON: %v", err)
			return
		}

		// Debug print the loaded emailData
		log.Printf("Loaded EmailData: %+v\n", emailData)

		// Execute the template with the data
		if err := tmpl.Execute(w, emailData); err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
			log.Fatalf("Error rendering template: %v", err)
			return
		}
	})

	// Start the server on port 8080
	log.Println("Server starting on http://localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
