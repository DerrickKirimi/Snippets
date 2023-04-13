package main
import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)


func (app *application) newTemplateData (r *http.Request) *templateData {
	return &templateData {
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	//Use output function and set frame depth to 2
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int){
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render (w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page] //Find the page in the map
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)


	//write to buffer rather than straight to http.ResponseWriter
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//if template written to buffers without errors, write http status code to http.ResponseWriter
	w.WriteHeader(status)

	buf.WriteTo(w)
}

