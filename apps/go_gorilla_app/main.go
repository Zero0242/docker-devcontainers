package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Zero0242/go_gorilla_app/config"
	"github.com/Zero0242/go_gorilla_app/db"
	"github.com/Zero0242/go_gorilla_app/middleware"
	"github.com/Zero0242/go_gorilla_app/models"
	"github.com/Zero0242/go_gorilla_app/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Lectura de variables
	config.GetEnvVars()

	// Config de base de datos
	db.Connect()
	db.DB.AutoMigrate(
		models.User{},
		models.Task{},
	)

	// Server
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	// API /users
	r.HandleFunc("/api/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/api/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	// API /tasks
	r.HandleFunc("/api/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/api/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/api/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", routes.UpdateTaskHandler).Methods("PATCH")
	r.HandleFunc("/api/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	// Middleware
	r.Use(middleware.LogMiddleware)

	// Listener
	port := ":" + strconv.Itoa(config.EnvMap.PORT)
	log.Println("Corriendo en el puerto " + port)
	http.ListenAndServe(port, r)
}
