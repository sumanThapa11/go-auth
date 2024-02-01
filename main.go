package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sumanThapa11/go-auth/cmd/routes"
	"github.com/sumanThapa11/go-auth/cmd/services"
)

func main() {
	l := log.New(os.Stdout, "auth-api", log.LstdFlags)
	db := services.GetDBConnection()

	ah := routes.NewAuth(l, db)

	r := mux.NewRouter()

	getRouter := r.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ah.SignUp)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server running on port 8080")
	log.Fatal(srv.ListenAndServe())

}
