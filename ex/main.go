package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleHome)

	http.ListenAndServe("localhost:4242", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {

	fmt.Println("connected to server")
	// return "ds";

}
