package main

import (
	"net/http"

	"github.com/fabiofns/api-ramada/Api/functions"

	"github.com/akrylysov/algnhsa"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("index"))
}

func addHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(functions.validateToken()))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)

	algnhsa.ListenAndServe(http.DefaultServeMux, nil)
}
