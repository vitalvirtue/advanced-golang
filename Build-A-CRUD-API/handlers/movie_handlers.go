package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/vitalvirtue/advanced-golang/Build-A-CRUD-API/models"
	"github.com/gorilla/mux"
	"strconv"
)

var movies []models.Movie

// Get all movies
func GetMovies(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get movie by ID
func GetMovie(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set the response header to application/json type to return JSON response to the client request URL /movies/{id} 
	params := mux.Vars(r) // Get params from request URL and return a map of string keys and string values like this: map["id"] = "1"
	id, _ := strconv.Atoi(params["id"]) // Convert the string value of the "id" key in the params map to an integer value

	for _, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie) // Return the movie object if the movie ID matches the ID in the request URL /movies/{id}
			return
		}
	}

	json.NewEncoder(w).Encode(&models.Movie{}) // Return an empty movie object if no movie with the specified ID is found in movies array
}

// Create a new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the JSON request body and store it in the movie object

	movie.ID = len(movies) + 1 // Set the movie ID to the length of the movies array + 1
	movies = append(movies, movie) // Append the movie object to the movies array
	json.NewEncoder(w).Encode(movie) // Return the movie object in the response
}

// Update a movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params from the request URL and return a map of string keys and string values
	id, _ := strconv.Atoi(params["id"]) // Convert the string value of the "id" key in the params map to an integer value

	var movie models.Movie // Create a movie object to store the updated movie object
	_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the JSON request body and store it in the movie object
	movie.ID = id

	for i, item := range movies {
		if item.ID == id {
			movies[i] = movie
			json.NewEncoder(w).Encode(movie) // Return the updated movie object in the response if the movie ID matches the ID in the request URL /movies/{id}
			return
		}
	}

	json.NewEncoder(w).Encode(&models.Movie{}) // Return an empty movie object if no movie with the specified ID is found in movies array
}

// Delete a movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the params from the request URL and return a map of string keys and string values
	id, _ := strconv.Atoi(params["id"]) // Convert the string value of the "id" key in the params map to an integer value

	for i, item := range movies {
		if item.ID == id {
			movies = append(movies[:i], movies[i+1:]...) // Remove the movie object from the movies array
			break
		}
	}

	json.NewEncoder(w).Encode(movies) // Return the updated movies array in the response
}