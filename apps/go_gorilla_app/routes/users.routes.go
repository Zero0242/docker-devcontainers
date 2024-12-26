package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Zero0242/go_gorilla_app/db"
	"github.com/Zero0242/go_gorilla_app/models"
	"github.com/gorilla/mux"
)

// Crea un usuario -.-
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err = createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("User Created"))
}

// Borra el usuario mediante su param
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User with id " + params["id"] + " not found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}

// Obtiene un usuario por id
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User with id " + params["id"] + " not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

// Obtiene un listado de usuarios
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}
