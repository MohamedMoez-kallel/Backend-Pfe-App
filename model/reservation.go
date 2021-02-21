package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Reservation struct {
	gorm.Model
	Heure_debut time.Time `json:"heure_debut"`
	Heure_fin   time.Time `json:"heure_fin"`
	User        User      `json:"-"`
	UserId      int       `json:"user_id"`
	SalleId     int       `json:"salle_id"`
}
