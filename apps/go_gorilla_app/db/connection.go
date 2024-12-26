package db

import (
	"log"
	"strconv"
	"strings"

	"github.com/Zero0242/go_gorilla_app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var DSN strings.Builder
	DSN.WriteString("host=" + config.EnvMap.DATABASE_HOST + " ")
	DSN.WriteString("user=" + config.EnvMap.DATABASE_USER + " ")
	DSN.WriteString("password=" + config.EnvMap.DATABASE_PASS + " ")
	DSN.WriteString("dbname=" + config.EnvMap.DATABASE_NAME + " ")
	DSN.WriteString("port=" + strconv.Itoa(config.EnvMap.DATABASE_PORT))
	var error error
	DB, error = gorm.Open(postgres.Open(DSN.String()), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connected")
	}
}
