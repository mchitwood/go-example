package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

// Movie Struct
type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", handleMovies).Methods("GET")
	http.ListenAndServe(os.Getenv("OPENSHIFT_GO_IP") + ":" + os.Getenv("OPENSHIFT_GO_PORT"), router)
}

func handleMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var movies = map[string]*Movie{
		"tt0076759": &Movie{Title: "Chucky's Revenge: The Never-Ending Cloud Party", Rating: "8.7", Year: "1977"},
		"tt0082971": &Movie{Title: "Eric The Red: The Search For a PasS", Rating: "8.6", Year: "1981"},
		"tt0086452": &Movie{Title: "Where in the World is Peter Springsteen?", Rating: "10.0", Year: "2016"},
	}

	outgoingJSON, error := json.Marshal(movies)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}