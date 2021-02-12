package model

import (
	"net/http"
	u "rh-projet/utils"
	"time"
)

type Evenement struct {
	Id                   int
	Titre                string `json:"titre"`
	Date_debut_evenement string `json:"date_debut_evn"`
	Date_fin_evenement   string `json:"date_fin_evn"`
	Emplacement          string `json:"emplacement"`
	Description          string `json:"description"`
	Photo                string `json:"photo"`
	//User                 User
	UserId       int    `json:"user_id"`
	Participants []User `gorm:"many2many:evenement_users;"`
}

func (evenement *Evenement) AjouterEvenement(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	GetDB().Create(evenement)
	response := u.Message("Evenement ajouter")
	return response
}

func AfficherEvenement() []*Evenement {

	evenement := []*Evenement{}
	err := GetDB().Table("evenements").Find(&evenement).Error
	if err != nil {
		return nil
	}
	return evenement
}

func DeletEvenement(id int) *Evenement {

	var evenement Evenement
	GetDB().First(&evenement, id)
	db.Delete(&evenement)
	return &evenement
}

func ModifierEvenement(evenement *Evenement) *Evenement {

	db.Save(&evenement)
	return evenement
}

func RechercheEvenement(title string) []*Evenement {

	evenement := []*Evenement{}
	GetDB().Table("evenements").Where("titre LIKE ? or emplacement LIKE ?  ", title, title).Find(&evenement)
	return evenement
}

func AfficherEvenementDate() []*Evenement {

	evenement := []*Evenement{}
	err := GetDB().Table("evenements").Where("Date_debut_evenement >= ?", time.Now()).Find(&evenement).Error
	if err != nil {
		return nil
	}
	return evenement
}
func AfficherParEve(id int) *Evenement {
	var evenement Evenement
	GetDB().Table("evenements").Where("id = ?", id).Find(&evenement)
	return &evenement
}
func AfficherUserEve(user_id int) []*Evenement {
	evenement := []*Evenement{}
	GetDB().Table("evenements").Where(" user_id =?", user_id).Find(&evenement)
	return evenement
}
func ReserverEvenement(participants []User, evenementID int) *Evenement {
	var evenement Evenement
	evenement.Id = evenementID

	if err := db.First(&evenement).Error; err != nil {
		panic(err)
	}
	db.Model(&evenement).Association("Participants").Append(participants)
	return &evenement

}
