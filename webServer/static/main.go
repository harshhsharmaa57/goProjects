package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "form.html")
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST req successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Address - %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
