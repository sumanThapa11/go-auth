package models

import (
	"encoding/json"
	"io"

	"gorm.io/gorm"
)

type USER struct {
	gorm.Model
	Name     string `gorm:"serializer:json"`
	Email    string `gorm:"serializer:json"`
	Password string `gorm:"serializer:json"`
}

func (u *USER) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}
