package main 

import (
	"log"
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippets"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Here lies your snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	if (r.Method != http.MethodPost){
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Let there be Light..."))
}

func main () {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err) 
}