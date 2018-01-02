package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Message struct {
	Text string
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	var msg string = "Hello.  My name is Super Snitch!  I check for plagiarism in
source code.";
	fmt.Fprintf(w, msg, r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request) {
	var msg string = "Hello.  My name is Super Snitch!";
	m := Message{
	fmt.Fprintf(w, msg, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/collection/:id", getCollection)
	http.HandleFunc("//:id", handler)
	http.ListenAndServe(":8080", nil)
}
