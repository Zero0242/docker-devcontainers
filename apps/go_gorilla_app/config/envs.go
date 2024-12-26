package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	PORT          int
	DATABASE_PORT int
	DATABASE_HOST string
	DATABASE_USER string
	DATABASE_PASS string
	DATABASE_NAME string
}

var EnvMap EnvVars

// GetEnvVars loads environment variables from a .env file and populates the EnvMap with these values.
// It uses the godotenv package to load the .env file and the os package to retrieve the environment variables.
// If the .env file cannot be loaded, the function logs a fatal error and terminates the program.
// The following environment variables are expected:
// - PORT: The port number on which the application will run.
// - DATABASE_PORT: The port number on which the database server is running.
// - DATABASE_HOST: The hostname or IP address of the database server.
// - DATABASE_USER: The username for connecting to the database.
// - DATABASE_PASS: The password for connecting to the database.
// - DATABASE_NAME: The name of the database to connect to.
// The function logs a message indicating that the environment configurations have been loaded successfully.
func GetEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	EnvMap = EnvVars{
		PORT:          port,
		DATABASE_PORT: dbPort,
		DATABASE_HOST: os.Getenv("DATABASE_HOST"),
		DATABASE_USER: os.Getenv("DATABASE_USER"),
		DATABASE_PASS: os.Getenv("DATABASE_PASS"),
		DATABASE_NAME: os.Getenv("DATABASE_NAME"),
	}

	log.Println("Cargada las configuraciones de entorno")
}
