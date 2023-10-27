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

	server := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	//Start server
	infoLog.Printf("Starting server on port %s", *addr)
	err := server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}
