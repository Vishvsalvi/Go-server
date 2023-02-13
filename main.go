package main

import (

	"net/http"
	"fmt"
	"strconv"
	"github.com/gorilla/mux"
	"math/rand"
	"log"
	"encoding/json"

)

type Movie struct {
	ID string `json:"id"` // The data in json will be represented as {id: someNumberString}
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") // Sets the header
	json.NewEncoder(w).Encode(movies) // Encodes the data into json 
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Accessing the URL params from the request

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...) // Exclude the movie from the slice
			break
		}

	}

}


func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" ,"application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item); // This line returns the data back to the client
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // Decodes the entire response body and stores it into the movie variable
	movie.ID = strconv.Itoa(rand.Intn(10000000)) // Randon no from 1 to 10 million
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) // Sends the info back to the client about the movie
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func main() {
	r := mux.NewRouter() // Creates a router instance
	movies = append(movies, Movie{ID: "1", Isbn: "456132", Title: "Black Panther",Director: &Director{Firstname: "Chadwick", Lastname: "Boseman"} })

	// The specific routes are handled by various functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000" ,r))

}