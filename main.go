package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"fmt"

	asciiart "ascii-art-web/ascii-art"
)

func main() {
	// Define HTTP handlers
	http.HandleFunc("/", Index)
	http.HandleFunc("/ascii-art", HandleASCIIArt)

	// Serve static files (uncomment and adjust as needed)
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Index handles requests to "/" and "/Home"
func Index(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/Home":
		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/index.html")
		}
	case "/about", "/About":
		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/about.html")
		}
	default:
		if r.Method == http.MethodGet {
			http.NotFound(w, r)
		}
	}
}

// HandleASCIIArt handles requests to "/ascii-art"
func HandleASCIIArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	bannerFile := r.FormValue("banner")
	input := r.FormValue("input")

	if len(bannerFile) == 0 || len(input) == 0 {
		http.Error(w, "Missing banner or input", http.StatusBadRequest)
		return
	}

	//if the specified banner file is missing return status code 500
	_, error := os.Stat(bannerFile)
	if os.IsNotExist(error){
		errorMsg := fmt.Sprintf("The banner file %s is missing", bannerFile)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}
	

	if containsIllegalCharacters(input) {
		http.Error(w, "Input contains illegal characters: Status 400: Bad Request", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	ascii := asciiart.Art(input, asciiart.BannerArt(asciiart.FileName(bannerFile)))
	tmpl.Execute(w, ascii.String())
}

// serveTemplate loads and executes a template file
func serveTemplate(w http.ResponseWriter, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	tmpl := template.Must(template.ParseFiles(filename))
	tmpl.Execute(w, nil)
}

// Function to check for illegal characters in input
func containsIllegalCharacters(input string) bool {
	// Regular expression to match non-printable ASCII characters
	illegalCharRegex := regexp.MustCompile(`[^\x00-\x7F]`)
	return illegalCharRegex.MatchString(input)
}
