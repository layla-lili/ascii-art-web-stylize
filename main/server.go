package main

import (
	"asciiart"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var templates *template.Template
var text string

func init() {
	// Load templates
	templates = template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/404.html",
		"templates/400.html",
		"templates/500.html",
	))
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Handler for the root URL ("/")
	http.HandleFunc("/", homeHandler)
	// Handler for the "/ascii-art" URL
	http.HandleFunc("/ascii-art", asciiArtHandler)
	// Handle not found: 404 page 
	http.HandleFunc("/404", NotFoundHandler)
	// Handle Bad Request : 400
	http.HandleFunc("/400", BadRequestHandler)
	// Handle Server error: 500
	http.HandleFunc("/500", internalServerErrorHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// homeHandler handles GET requests to the root URL ("/").
// It renders the index.html template.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
	}
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			internalServerErrorHandler(w, r)
			return
		}
	}
}

var textxolor = "#EE1A1A"

// asciiArtHandler handles POST requests to the "/ascii-art" URL.
// It generates ASCII art based on the input text and selected banner,
// and renders the index.html template with the generated ASCII art.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		text = r.FormValue("text")
		banner := r.FormValue("banner")
		textxolor = r.FormValue("textcolor")
		asciiArt := generateAsciiArt(text, banner, w, r)
		notENg := false
		for _, ch := range text {
			if ch > 127 || ch < 32 && r.URL.Path == "/ascii-art" {
				if ch != 10 && ch != 13 {
					notENg = true
				}
			}
		}
		if notENg {
			BadRequestHandler(w, r)
			return
		} else {

			if len(asciiArt) > 0 {
				err := templates.ExecuteTemplate(w, "index.html", struct {
					Text      string
					Banner    string
					AsciiArt  string
					TextColor string
				}{
					Text:      text,
					Banner:    banner,
					AsciiArt:  asciiArt,
					TextColor: textxolor,
				})
				if err != nil {
					internalServerErrorHandler(w, r)
					return
				}
			}

		}
	}
}
func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	err := templates.ExecuteTemplate(w, "400.html", nil)
	if err != nil {
		internalServerErrorHandler(w, r)
		return
	}
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := templates.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		internalServerErrorHandler(w, r)
		return
	}
}
func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	templates.ExecuteTemplate(w, "500.html", nil)
}

func generateAsciiArt(text, banner string, w http.ResponseWriter, r *http.Request) string {
	// Implement your ASCII art generation logic based on the selected banner
	// Here's a simple example for the three banners mentioned
	result := ""
	lines := strings.Split(text, "\n")
	file, err := os.Open(banner + ".txt")
	if err != nil {
		internalServerErrorHandler(w, r)
		//fmt.Println(err.Error())
		//os.Exit(1)
	} else {
		defer file.Close()

		for _, line := range lines {
			//check if the line empty
			if line != "" {
				//call the ReadLine function to print Ascii art for the current line
				result = result + asciiart.ReadLine(line, file)
			}
		}
	}
	return result
}
