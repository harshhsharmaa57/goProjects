package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movies struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movies

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}





func main() {
	r := mux.NewRouter()

	movies = append(movies, Movies{ID: "1", Isbn: "5344", Title: "HarryPotter", Director: &Director{FirstName: "JK", LastName: "Rowling"}})
	movies = append(movies, Movies{ID: "2", Isbn: "5435", Title: "Spiderman", Director: &Director{FirstName: "Tom", LastName: "Holland"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server starting at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}