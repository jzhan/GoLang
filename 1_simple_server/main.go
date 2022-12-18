package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request: ", r.URL.Path)

		if r.URL.Path == "/" {
			fmt.Println("URL: ", r.URL.Path)

			n, err := fmt.Fprintf(w, "<h1>hello world</h1>")

			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("Number of bytes written: %d\n", n)
		}
	})

	http.ListenAndServe(":8080", nil)
}
