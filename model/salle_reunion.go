package model

import (
	"fmt"
	"net/http"
	u "rh-projet/utils"
	"time"
)

type Salle_reunion struct {
	Id          int
	Num_salle   int       `json:"num_salle"`
	Date_debut  string    `json:"date_debut"`
	Heure_debut time.Time `json:"heure_debut"`
	Heure_fin   time.Time `json:"heure_fin"`
	User        User
	UserId      int `json:"user_id"`
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

func ReserverSalle(user_id int, Date_debut string, heure_debut, heure_fin time.Time) *Salle_reunion {
	// var result int
	// salle := []*Salle_reunion{}
	// db.Table("salle_reunions").Where("user_id=0").Count(&result)
	// err := GetDB().Table("salle_reunions").Where("user_id=0").Find(&salle).Error
	// if err != nil {
	// 	return nil
	// }
	// if result > 0 {
	// 	salle[0].UserId = user_id
	// 	salle[0].Date_debut = Date_debut
	// 	salle[0].Heure_debut = Heure_debut
	// 	salle[0].Heure_fin = Heure_fin

	// 	db.Table("salle_reunions").Save(&salle[0])
	// }
	// return salle[0]
	salles := []Salle_reunion{}
	err := GetDB().Model(&Salle_reunion{}).Where("user_id = ?", user_id).Find(salles).Error
	if err != nil {
		fmt.Println(err)
	}
	if len(salles) <= 0 {
		return nil
	}
	var salle = salles[0]
	salle.UserId = user_id
	salle.Heure_debut = heure_debut
	salle.Heure_fin = heure_fin
	err = GetDB().Model(&Salle_reunion{}).Updates(&salle).Error
	if err != nil {
		fmt.Println(err)
	}
	return &salle
}
