package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Scores struct {
	Score map[string]int
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	if filename == "" {
		filename = "index.html"
	}

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(w, err.Error(), 404)
	} else {
		w.Write(f)
	}
}
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fruit := vars["fruit"]

	fmt.Println(fruit)

}

func GetResultsHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", FileHandler)
	r.HandleFunc("/{filename}", FileHandler)

	r.HandleFunc("/submit/{fruit}", SubmitHandler)
	r.HandleFunc("/getResults", GetResultsHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
