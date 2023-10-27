package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Struct to hold our application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//Add command line flag for addr value
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	//Specify log levels
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//initialise struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//Multiplexer for routing
	mux := http.NewServeMux()

	//File server that points to the static file directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/snippet/view", app.viewSnippetsHandler)
	mux.HandleFunc("/snippet/create", app.createSnippetHandler)

	server := &http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: errorLog,
	}

	//Start server
	infoLog.Printf("Starting server on port %s", *addr)
	err := server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}
