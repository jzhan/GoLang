package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><a href=\"about\">home</a><h1>")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1><a href=\"home\">about</a><h1>")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(PORT, nil)
}
