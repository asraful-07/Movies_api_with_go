package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found")
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
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

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Movie deleted")
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1111", Title: "Inception", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1211", Title: "The Godfather", Director: &Director{Firstname: "Francis", Lastname: "Coppola"}})
	movies = append(movies, Movie{ID: "3", Isbn: "1311", Title: "Interstellar", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "4", Isbn: "1411", Title: "Parasite", Director: &Director{Firstname: "Bong", Lastname: "Joon-ho"}})
	movies = append(movies, Movie{ID: "5", Isbn: "1511", Title: "The Dark Knight", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "6", Isbn: "1611", Title: "Pulp Fiction", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// CORS Middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow your frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
