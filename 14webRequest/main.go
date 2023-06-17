package main

import (
	"fmt"
	"net/http"

	"github.com/surya7765/routes"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Print(routes.Multiply())
	http.ListenAndServe(":8080", nil)
}
