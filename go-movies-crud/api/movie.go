package api

import (
	"encoding/json"
	"go-movies/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Movies []model.Movie
var Director []model.Director

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func GetDirector(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Director)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	newMovie := []model.Movie{}

	for _, x := range Movies {
		if x.ID != id {
			newMovie = append(newMovie, x)
		}
	}

	Movies = newMovie

	json.NewEncoder(w).Encode(map[string]string{"messsage": "The data has been deleted succesfully"})

}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	newData := []model.Movie{}
	for _, x := range Movies {
		if x.ID != id {
			newData = append(newData, x)
		}
	}

	Movies = newData

	json.NewEncoder(w).Encode(map[string]string{"message": "The selected data has been deletedI"})
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	selectedId := params["id"]

	// data := []model.Movie{}

	for _, x := range Movies {
		if x.ID == selectedId {
			json.NewEncoder(w).Encode(x)
			// data = append(data, x)
			break
		}
	}

	// Movies = data

	// json.NewEncoder(w).Encode(Movies)
}

// func CreateMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	var movie model.Movie

// 	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
// 		http.Error(w, "invalid request", http.StatusBadRequest)
// 	}
// 	_ = json.NewDecoder(r.Body).Decode(&movie)
// 	movie.ID = strconv.Itoa(len(Movies) + 1)
// 	Movies = append(Movies, movie)

// 	json.NewEncoder(w).Encode(Movies)
// }

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	movie.ID = strconv.Itoa(len(Movies) + 1)
	Movies = append(Movies, movie)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	selectedId := params["id"]

	var newData model.Movie

	if err := json.NewDecoder(r.Body).Decode(&newData); err != nil {
		http.Error(w, "Invaalid request", http.StatusBadRequest)
		return
	}

	update := false

	for i := range Movies {
		if Movies[i].ID == selectedId {
			Movies[i] = newData
			update = true
			break
		}
	}

	if !update {
		http.Error(w, "Failed to update the data", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(Movies)
}
