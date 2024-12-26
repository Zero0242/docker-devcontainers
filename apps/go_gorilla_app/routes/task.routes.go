package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Zero0242/go_gorilla_app/db"
	"github.com/Zero0242/go_gorilla_app/models"
	"github.com/gorilla/mux"
)

// Retorna la tarea
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task with id " + params["id"] + " not found"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

// Crea una nueva tarea
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	if task.UserID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User ID is required"))
		return
	}

	result := db.DB.Create(&task)
	err := result.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something went wrong :c " + err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

// Actualiza una tarea existente
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task with id " + params["id"] + " not found"))
		return
	}

	json.NewDecoder(r.Body).Decode(&task)
	db.DB.Save(&task)
	json.NewEncoder(w).Encode(&task)
}

// Elimina una tarea
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task with id " + params["id"] + " not found"))
		return
	}

	db.DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
