package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetBookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
}

func GetBookDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	selectedId := params["id"]
	print(selectedId)

}
