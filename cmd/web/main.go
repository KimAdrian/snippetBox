package main

import (
	"log"
	"net/http"
)

func main() {
	//Multiplexer for routing
	mux := http.NewServeMux()

	//File server that points to the static file directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/view", viewSnippetsHandler)
	mux.HandleFunc("/snippet/create", createSnippetHandler)

	//Start server
	log.Print("Starting server on port 8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalf("Error: %d", err)
	}

}
