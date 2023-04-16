package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"github.com/DerrickKirimi/Snippets/internal/models"
	"github.com/go-playground/form/v4"

	_ "github.com/go-sql-driver/mysql"
)

//application-wide dependencies definition
type application struct {
	errorLog 		*log.Logger
	infoLog  		*log.Logger
	snippets 		*models.SnippetModel
	templateCache	map[string]*template.Template
	formDecoder		*form.Decoder
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "web:espresso@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL data source name")

	//parse into addr variable. Terminate application on error.
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn) //pass in the dsn from the flags
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close() //defer a call to db.Close() so before main exits connection pool is closed

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	//dependencies initialisation
	app := &application{
		errorLog: 		errorLog,
		infoLog: 		infoLog,
		snippets: 		&models.SnippetModel{DB: db},
		templateCache:	templateCache,
		formDecoder: 	formDecoder,
	}

	

	//same network addr and routes as before + ErrorLog field set
	srv := &http.Server{
		Addr:		*addr,
		ErrorLog:	errorLog,
		Handler: 	app.routes(),

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