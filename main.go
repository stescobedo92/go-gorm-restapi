package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stescobedo92/go-gorm-restapi/db"
	"github.com/stescobedo92/go-gorm-restapi/routes"
)

func main() {
	db.DBConnection()

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
