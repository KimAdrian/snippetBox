package main

import (
	"log"
	"net/http"
)

func main() {
	//Multiplexers for routing
	mux := http.NewServeMux()
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
