package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	//Multiplexer for routing
	mux := http.NewServeMux()

	//File server that points to the static file directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/snippet/view", app.viewSnippetsHandler)
	mux.HandleFunc("/snippet/create", app.createSnippetHandler)

	return mux
}
