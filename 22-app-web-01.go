package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bonjour à tous!")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "À propos de nous")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact-nous")
}

func appWeb() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)

	fmt.Println("(http://localhost:8080) - Server running on port", port)
	http.ListenAndServe(port, nil)
}
