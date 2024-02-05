package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sumanThapa11/go-auth/cmd/models"
	"github.com/sumanThapa11/go-auth/cmd/routes"
	"github.com/sumanThapa11/go-auth/cmd/services"
)

func main() {
	l := log.New(os.Stdout, "auth-api", log.LstdFlags)
	db := services.GetDBConnection()

	user := models.USER{}

	db.AutoMigrate(&user)

	ah := routes.NewAuth(l, db)

	r := mux.NewRouter()

	getRouter := r.Methods("GET").Subrouter()
	getRouter.HandleFunc("/getUserById/{id:[0-9]+}", ah.GetUserById)
	getRouter.HandleFunc("/getUsersDetails", ah.GetAllUserDetails)

	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/signup", ah.SignUp)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server running on port 8080")
	log.Fatal(srv.ListenAndServe())

}
