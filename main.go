package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

	f, err := os.OpenFile("results.csv", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fruit := r.PostFormValue("fruit")

	_, err = f.WriteString(fruit + "\n")
	if err != nil {
		fmt.Println("error:", err)
	}

}

func GetResultsHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", FileHandler)
	r.HandleFunc("/{filename}", FileHandler)

	r.HandleFunc("/submit/", SubmitHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
