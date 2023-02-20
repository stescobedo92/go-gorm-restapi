package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, rquest *http.Request) {
		writer.Write([]byte("Hello world"))
	})

	http.ListenAndServe(":3000", router)
}
