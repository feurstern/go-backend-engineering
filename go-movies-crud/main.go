package main

import (
	"fmt"
	"go-movies/api"
	"go-movies/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	api.Movies = append(api.Movies, model.Movie{ID: "1", Isbn: "NSDAP-12312", Title: "Mein kampf", Director: &model.Director{Firstname: "Adolf", Lastname: "Hitler"}})
	api.Movies = append(api.Movies, model.Movie{ID: "2", Isbn: "ac12121", Title: "The lord of the rings", Director: &model.Director{Firstname: "Hatssune", Lastname: "Miku"}})

	api.Director = append(api.Director, model.Director{Firstname: "Adolf", Lastname: "Hilter"})

	r := mux.NewRouter()

	r.HandleFunc("/movies", api.GetMovies).Methods("GET")

	r.HandleFunc("/director", api.GetDirector).Methods("GET")
	// r.HandleFunc("/movies/{id}", api.GetMovie).Methods("GET")
	r.HandleFunc("/movies/create", api.CreateMovie).Methods("POST")
	// r.HandleFunc("/movies/update/{id}", updateMove).Methods("POST")
	r.HandleFunc("/movies/delete/{id}", api.DeleteMovie).Methods("POST")

	fmt.Printf("starting server at port: 4242 \n")
	log.Fatal(http.ListenAndServe(":4242", r))

}
