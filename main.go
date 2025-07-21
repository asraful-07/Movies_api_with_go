// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/rs/cors"
// )

// type Movie struct {
// 	ID       string    `json:"id"`
// 	Isbn     string    `json:"isbn"`
// 	Title    string    `json:"title"`
// 	Image    string    `json:"image"`
// 	Description string  `json:"description"`
// 	Director *Director `json:"director"`
// }

// type Director struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// var movies []Movie

// func getMovies(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(movies)
// }

// func getMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, item := range movies {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode("Movie not found")
// }

// func createMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var movie Movie
// 	_ = json.NewDecoder(r.Body).Decode(&movie)
// 	movies = append(movies, movie)
// 	json.NewEncoder(w).Encode(movie)
// }

// func updateMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			var movie Movie
// 			_ = json.NewDecoder(r.Body).Decode(&movie)
// 			movie.ID = params["id"]
// 			movies = append(movies, movie)
// 			json.NewEncoder(w).Encode(movie)
// 			return
// 		}
// 	}
// }

// func deleteMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode("Movie deleted")
// }

// func main() {
// 	r := mux.NewRouter()

// 	movies = append(movies, Movie{ ID: "1", Isbn: "1111", Title: "Baahubali", Image: "https://i.ibb.co/ZRyD4qdK/baahubali.jpg", Description:"In the ancient kingdom of Mahishmati, a young man named Shivudu discovers his royal lineage and rises to challenge the tyrant who wronged his family. An epic saga of courage, betrayal, and destiny.", Director: &Director{ Firstname: "S. S.",Lastname: "Rajamouli"}})
//     movies = append(movies, Movie{ ID: "2", Isbn: "1211", Title: "3 Idiots", Image: "https://i.ibb.co/srk36ts/3Idiots.jpg", Description: "Three engineering students navigate their way through the pressures of an unforgiving academic system while discovering the true meaning of friendship, education, and chasing dreams.", Director: &Director{	Firstname: "Rajkumar", Lastname: "Hirani"}})
//     movies = append(movies, Movie{ ID: "3", Isbn: "1311", Title: "Avengers", Image: "https://i.ibb.co/VY6VZjYk/avengers.jpg", Description:"When global peace is threatened by Loki and his alien army, Earth's mightiest heroes Iron Man, Thor, Captain America, Hulk, Black Widow, and Hawkeye must unite to defend the planet in an explosive battle for humanity.", Director: &Director{ Firstname: "Anthony", Lastname: "Russo"}})
// 	movies = append(movies, Movie{ ID: "4", Isbn: "1411", Title: "Spider-Man", Image: "https://i.ibb.co/SXWV7j6V/spider.jpg", Description: "After being bitten by a radioactive spider, teenager Peter Parker gains spider-like powers and must balance his personal life with his responsibilities as Spider-Man to protect New York from formidable villains.", Director: &Director{ Firstname: "Jon",	Lastname: "Watts"}})
// 	movies = append(movies, Movie{ID: "5", Isbn: "1511", Title: "Squid Game", Image: "https://i.ibb.co/p60Mt3x6/squid-game.jpg", Description: "Struggling financially, hundreds of players accept a mysterious invitation to compete in deadly children's games. The prize? Billions of won. The cost? Their lives. A chilling survival drama that explores human desperation.", Director: &Director{ Firstname: "Hwang", Lastname: "Dong-hyuk"}})
// 	movies = append(movies, Movie{ID: "6", Isbn: "1611", Title: "Hera Pheri", Image: "https://i.ibb.co/6RBNqDTn/hera-pheri.jpg", Description: "A hilarious rollercoaster of misadventures follows three broke men — a landlord, a tenant, and an out-of-work man — who accidentally become part of a kidnapping plot, leading to laugh-out-loud chaos and confusion.", Director: &Director{ Firstname: "Priyadarshan", Lastname: "Tarantino"}})

// 	r.HandleFunc("/movies", getMovies).Methods("GET")
// 	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
// 	r.HandleFunc("/movies", createMovie).Methods("POST")
// 	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
// 	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

// 	// CORS Middleware
// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"http://localhost:5173"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Content-Type"},
// 		AllowCredentials: true,
// 	})

// 	handler := c.Handler(r)

// 	fmt.Println("Server started at :8080")
// 	log.Fatal(http.ListenAndServe(":8080", handler))
// }

package main

import "fmt" 

func main()  {
	fmt.Println("hello world")
}
