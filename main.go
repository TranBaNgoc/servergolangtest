package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
