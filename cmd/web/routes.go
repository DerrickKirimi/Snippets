package main 
import (
	"net/http"
)

func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static", http.StripPrefix("/static", fileServer))

	mux.Handle("/", app.home)
	mux.Handle("/snippet/view", app.snippetView)
	mux.Handle("/snipet/create", app.snippetCreate)

	return mux
}