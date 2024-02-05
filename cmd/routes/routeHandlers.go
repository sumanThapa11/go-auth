package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sumanThapa11/go-auth/cmd/models"
	"gorm.io/gorm"
)

type Auth struct {
	l  *log.Logger
	db *gorm.DB
}

func NewAuth(l *log.Logger, db *gorm.DB) *Auth {
	return &Auth{l, db}
}

func (a *Auth) GetUserById(w http.ResponseWriter, r *http.Request) {
	a.l.Println("Handling get single user by id")

	vars := mux.Vars(r)

	var user models.USER

	userId, err := strconv.Atoi(vars["id"])

	if err != nil {
		a.l.Println("error while converting user id to int")
		return
	}

	result := a.db.First(&user, userId)

	if result.Error != nil {
		a.l.Fatalf("Error while getting user details, %d", userId)
		a.l.Println(result.Error)
		return
	}

	userDetails, _ := json.Marshal(user)

	a.l.Println(string(userDetails))

	w.Write([]byte(string(userDetails)))
}

func (a *Auth) GetAllUserDetails(w http.ResponseWriter, r *http.Request) {
	a.l.Println("Handling get all user details request")

	var users []models.USER

	result := a.db.Find(&users)

	if result.Error != nil {
		a.l.Fatal("Error while fetching user data from db")
		a.l.Panicln(result.Error)
		return
	}

	userDetails, _ := json.Marshal(users)

	a.l.Println(string(userDetails))
	// resp := strconv.Itoa(int(result.RowsAffected))

	w.Write([]byte(string(userDetails)))
	a.l.Println("Successfully fetched user data")
}

func (a *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	a.l.Println("Handling post request")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var user models.USER

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		a.l.Println(err)
		return
	}

	result := a.db.Create(&user)

	if result.Error != nil {
		a.l.Fatal("Error while inserting user data to db")
		a.l.Panicln(result.Error)
		return
	}

	resp := strconv.Itoa(int(result.RowsAffected))

	w.Write([]byte(resp))
	a.l.Println("Successfully created user data")
}
