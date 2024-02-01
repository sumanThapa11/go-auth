package routes

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Auth struct {
	l  *log.Logger
	db *gorm.DB
}

func NewAuth(l *log.Logger, db *gorm.DB) *Auth {
	return &Auth{l, db}
}

func (a *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	a.l.Println("Handling post request")
}
