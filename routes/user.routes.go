package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stescobedo92/go-gorm-restapi/db"
	"github.com/stescobedo92/go-gorm-restapi/models"
)

func GetUsersHandler(writer http.ResponseWriter, rquest *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(writer).Encode(&users)
}

func GetUserHandler(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	params := mux.Vars(request)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("User not found"))
	}

	json.NewEncoder(writer).Encode(params)
}

func PostUsersHandler(writer http.ResponseWriter, rquest *http.Request) {
	var user models.User
	json.NewDecoder(rquest.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
	}

	json.NewEncoder(writer).Encode(&user)
}

func DeleteUsersHandler(writer http.ResponseWriter, rquest *http.Request) {
	writer.Write([]byte("GetUsersHandler"))
}
