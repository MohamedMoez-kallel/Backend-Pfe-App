package model

import (
	"errors"
	"net/http"
	u "rh-projet/utils"
	"time"

	"github.com/jinzhu/gorm"
)

type Salle_reunion struct {
	gorm.Model
	NomSalle     string        `json:"nom_salle"`
	Reservations []Reservation `json:"reservations" gorm:"foreignKey:SalleId"`
}

func (salle_reunion *Salle_reunion) InsertSalleReunion(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	GetDB().Create(salle_reunion)
	response := u.Message("Salle réunion ajouter avec succès")
	return response
}

func AfficherSalle() []*Salle_reunion {
	salle := []*Salle_reunion{}
	err := GetDB().Table("salle_reunions").Find(&salle).Error
	if err != nil {

	}
	return salle
}
func SupprimerSalle(id int) *Salle_reunion {
	var salle Salle_reunion
	GetDB().First(&salle, id)
	db.Delete(&salle)
	return &salle
}

func ModifierSalle(salle *Salle_reunion) *Salle_reunion {
	db.Save(&salle)
	return salle
}

func ReserverSalle(salleId, userId int, Date_debut string, heure_debut, heure_fin time.Time) (*Salle_reunion, error) {
	var salle Salle_reunion
	err := GetDB().Model(&salle).Where("id = ?", salleId).First(&salle).Error
	if err != nil {
		return nil, err
	}
	err = GetDB().Model(&salle).Related(&salle.Reservations, "Reservations").Error
	if err != nil {
		return nil, err
	}
	for _, reservation := range salle.Reservations {
		if inTimeSpan(reservation.Heure_debut, reservation.Heure_fin, heure_debut) {
			return nil, errors.New("there is a reserved session! try with another time slot")
		}
	}
	err = GetDB().Model(&salle).Association("Reservations").Append([]Reservation{
		{
			Heure_debut: heure_debut,
			Heure_fin:   heure_fin,
			UserId:      userId,
		},
	}).Error
	if err != nil {
		return nil, err
	}
	return &salle, nil
}

func inTimeSpan(start, end, check time.Time) bool {
    return check.After(start) && check.Before(end)
}