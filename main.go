package main

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

type AsciiArt struct {
	Result string
}

type ErrorPageData struct {
	Title      string // Title of the HTML page
	StatusCode int    // HTTP status code
	Message    string // Error message
}

var exportContent string
var errorPageTemplate *template.Template

func main() {
	launchServer()
}

func launchServer() {
	// Define the directory to serve files from
	dir := http.Dir(".")

	// Create a file server instance
	fileServer := http.FileServer(dir)

	// Register the file server handler to serve files from the current directory
	http.Handle("/", fileServer)

	// Register handler for /ascii-art endpoint
	http.HandleFunc("/ascii-art", printHandler)
	// Register handler for /export endpoint
	http.HandleFunc("/export", exportHandler)

	// Load and parse the error page HTML file
	errorPageTemplate = template.Must(template.ParseFiles("templates/errorPage.html"))

	// Start the server on localhost port 8080
	addr := "localhost:8080"
	fmt.Print("Server listening on http://" + addr + "\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func printHandler(w http.ResponseWriter, r *http.Request) {
	// Check if request method is POST
	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}

		// Access form data
		asciiInput := r.Form.Get("ascii-input")
		artType := r.Form.Get("art-type")

		// Do something with the data
		ascii, err := returnAscii(asciiInput, artType)
		if err != nil {
			var statusCode int
			switch err.Error() {
			case "no input", "invalid character":
				statusCode = http.StatusBadRequest
			case "file corrupted", "banner file not found":
				statusCode = http.StatusInternalServerError
			default:
				statusCode = http.StatusInternalServerError
			}

			// Create ErrorPageData for the error
			errorData := createErrorPageData(statusCode, err)

			// Execute the error page template with the error data
			w.WriteHeader(statusCode)
			errorPageTemplate.Execute(w, errorData)
			return
		}

		// If no error, continue with rendering the ASCII art
		data := AsciiArt{
			Result: ascii,
		}

		// Parse HTML template
		tmpl, err := template.ParseFiles("templates/ascii-art.html")
		if err != nil {
			handleError(w, http.StatusNotFound, err)
			return
		}

		// Execute template with data
		err = tmpl.Execute(w, data)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}
		exportContent = ascii
	} else {
		// Create ErrorPageData for the error
		errorData := createErrorPageData(http.StatusMethodNotAllowed, errors.New("method not allowed"))

		// Execute the error page template with the error data
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorPageTemplate.Execute(w, errorData)
	}
}

func exportHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response headers
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(exportContent)))

	// Write the content to the response body
	w.Write([]byte(exportContent))
}
