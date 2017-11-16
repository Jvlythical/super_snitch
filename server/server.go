package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var msg string = "Hello.  My name is Super Snitch!";
	fmt.Fprintf(w, msg, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/collection/:id", getCollection)
	http.HandleFunc("//:id", handler)
	http.ListenAndServe(":8080", nil)
}
