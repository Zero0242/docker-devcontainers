package routes

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Handled the home path"))
	w.Write([]byte("Hello from the gorilla creature"))
}
