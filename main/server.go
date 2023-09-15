package main

import (
	"asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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

	// Handle not found: 404
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
		err := templates.ExecuteTemplate(w, "304.html", nil)
		if err != nil {
			internalServerErrorHandler(w,r)
			return
		}
	}
	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			internalServerErrorHandler(w,r)
			return
		}
	}
}

// asciiArtHandler handles POST requests to the "/ascii-art" URL.
// It generates ASCII art based on the input text and selected banner,
// and renders the index.html template with the generated ASCII art.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	textxolor:="#EE1A1A"
	if r.Method == http.MethodPost {
		text = r.FormValue("text")
		banner := r.FormValue("banner")
		textxolor = r.FormValue("textcolor")
		fmt.Println("textxolor",textxolor)
		asciiArt := generateAsciiArt(text, banner)
		notENg := false
		for _, ch := range text {
			if ch > 127 || ch < 32 && r.URL.Path == "/ascii-art" {
				notENg = true
			}
		}
		if notENg {
			err := templates.ExecuteTemplate(w, "400.html", nil)
			if err != nil {
				internalServerErrorHandler(w,r)
				return
			}
			return
		} else {
			err := templates.ExecuteTemplate(w, "index.html", struct {
				Text     string
				Banner   string
				AsciiArt string
				TextColor string
			}{
				Text:     text,
				Banner:   banner,
				AsciiArt: asciiArt,
				TextColor: textxolor,

			})
			if err != nil {
				internalServerErrorHandler(w,r)
				return
			}
		}
	}

}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	templates.ExecuteTemplate(w, "400.html", nil)

}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.ExecuteTemplate(w, "404.html", nil)
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	templates.ExecuteTemplate(w, "500.html", nil)
}

func generateAsciiArt(text, banner string) string {
	// Implement your ASCII art generation logic based on the selected banner
	// Here's a simple example for the three banners mentioned
	file, err := os.Open(banner + ".txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	return asciiart.ReadLine(text, file)
}
