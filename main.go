package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	filename := "." + r.URL.Path

	if r.URL.Path == "/" {
		filename = "index.html"
	}
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		w.Write(f)
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", FileHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
