package main

import (
	"net/http"
	"strings"
	"text/template"
)

const PORT = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()

	if url == "/" {
		url = "/home.html"
	} else {
		// find period index position
		pos := strings.LastIndex(url, ".")

		var extension string = ""

		// try to get the file extension
		if pos != -1 {
			extension = url[pos:]
		}

		// if the file doesn't have extension, add .html
		if extension == "" {
			url = url + ".html"
		}
	}

	parsed_template, err := template.ParseFiles("./templates/" + url)

	if err != nil {
		parsed_template, err = template.ParseFiles("./templates/" + url + ".html")
	}

	if err == nil {
		parsed_template.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(PORT, nil)
}
