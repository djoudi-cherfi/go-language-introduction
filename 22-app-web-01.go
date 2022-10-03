package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func appWeb() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	fmt.Println("(http://localhost:8080) - Server running on port", port)
	http.ListenAndServe(port, nil)
}
