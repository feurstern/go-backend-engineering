package routes

import (
	"go-book/pkg/controllers"

	"github.com/gorilla/mux"
)

var BookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/list", controllers.GetBookList).Methods("GET")
}
