package main

import (
	"database/sql"
	"flag"
	"github.com/KimAdrian/snippetBox/internal/model"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

// Struct to hold our application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *model.SnippetModel
}

func main() {

	//Command line flags
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String(
		"dsn",
		"postgres://snippetbox:root@localhost:5432/snippetbox?sslmode=disable",
		"Postgres database datasource name")
	flag.Parse()

	//Specify log levels
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Open connection pool to database
	dbConn, err := sql.Open("postgres", *dsn)
	if err != nil {
		errorLog.Fatalf("Can't connect to database: %v\n", err)
	}

	defer dbConn.Close()

	//initialise struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &model.SnippetModel{DB: dbConn},
	}

	server := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	//Start server
	infoLog.Printf("Starting server on port %s", *addr)
	err = server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}
