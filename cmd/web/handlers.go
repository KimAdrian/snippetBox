package main

import (
	"fmt"
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

	_, err := writer.Write([]byte("Hello from snippetBox"))
	if err != nil {
		log.Fatal(err)
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
