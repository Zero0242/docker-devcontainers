package routes

import "net/http"

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("User Created"))

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Found!"))
}
