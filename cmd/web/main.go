package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//application-wide dependencies definition
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Http network address")

	//parse into addr variable. Terminate application on error.
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//dependencies initialisation
	app := application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	

	//same network addr and routes as before + ErrorLog field set
	srv := &http.Server{
		Addr:		*addr,
		ErrorLog:	errorLog,
		Handler: 	app.routes()

	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
