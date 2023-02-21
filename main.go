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
	router.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	router.HandleFunc("/users", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
