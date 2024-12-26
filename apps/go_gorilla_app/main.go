package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Zero0242/go_gorilla_app/config"
	"github.com/Zero0242/go_gorilla_app/db"
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

	// Listener
	port := ":" + strconv.Itoa(config.EnvMap.PORT)
	log.Println("Corriendo en el puerto" + port)
	http.ListenAndServe(port, r)
}
