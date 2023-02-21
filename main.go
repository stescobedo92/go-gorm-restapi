package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stescobedo92/go-gorm-restapi/db"
	"github.com/stescobedo92/go-gorm-restapi/models"
	"github.com/stescobedo92/go-gorm-restapi/routes"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
