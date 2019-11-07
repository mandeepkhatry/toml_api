package main

import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", keyHandler)
	r.HandleFunc("/{category}", categoryHandler)
}
