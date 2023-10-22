package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	//Restrict root pattern from being a fallback option
	//Send 404 response of url specified is not recognized and return
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	//Initialise slice containing path to templates
	templateFiles := []string{
		"./ui/html/templates/base.html",
		"./ui/html/templates/nav.html",
		"./ui/html/templates/pages/home.html",
	}
	//Read template file into template set
	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	//Write template content as the response body
	err = ts.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
	}
}

func viewSnippetsHandler(writer http.ResponseWriter, request *http.Request) {
	//Fetch id parameter from url and convert it to int
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	//If id cannot be converted or is less than 1 respond with 404 and return
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	fmt.Fprintf(writer, "Display a specific snippet with ID %d", id)
}

func createSnippetHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}

	writer.Write([]byte("Creating new snippet..."))
}
