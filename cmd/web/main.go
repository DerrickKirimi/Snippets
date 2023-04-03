package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//application-wide dependencies definition
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MYSQL data source name")


	//parse into addr variable. Terminate application on error.
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn) //pass in the dsn from the flags
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close() //defer a call to db.Close() so before main exits connection pool is closed
	//dependencies initialisation
	app := &application{
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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return  nil, err
	}
	return db, nil
}