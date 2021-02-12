package model

import (
	"fmt"
	"net/http"
	u "rh-projet/utils"
	"time"
)

type Formation struct {
	//gorm.Model

	Id              int
	Heure_formation string `json:"heure_formation"`
	Date_formation  string `json:"date_formation"`
	Titre           string `json:"titre"`
	Description     string `json:"description"`
	Formateur       string `json:"formateur"`
	Nb_place        int64  `json:"nb_place"`
	Type_formation  string `json:"type_formation"`
	Photo           string `json:"photo"`
	//User            User
	UserId       int    `json:"user_id"`
	Participants []User `gorm:"many2many:formation_users;"`
}

func (formation *Formation) InsertFormation(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	GetDB().Create(formation)
	response := u.Message("Formation ajouter avec succès")
	return response
}

func AfficherFormation() []*Formation {

	formation := []*Formation{}
	err := GetDB().Table("formations").Find(&formation).Error
	if err != nil {
		return nil
	}
	return formation
}

func SupprimerFormation(id int) *Formation {

	var formation Formation
	u := GetDB().First(&formation, id)
	fmt.Println(u)
	db.Delete(&formation)

	return &formation
}

func ModifierFormation(formation *Formation) *Formation {

	db.Save(&formation)
	return formation
}

func RechercheFormation(title string) []*Formation {

	formation := []*Formation{}
	GetDB().Table("formations").Where(" titre LIKE ? or type_formation LIKE ? ", title, title).Find(&formation)
	return formation
}

func AfficherFormationDate() []*Formation {

	formation := []*Formation{}
	err := GetDB().Table("formations").Where("date_formation >= ?", time.Now()).Find(&formation).Error
	if err != nil {
		return nil
	}
	return formation
}
func AfficherParFor(id int) *Formation {
	var formation Formation
	GetDB().Table("formations").Where("id = ?", id).Find(&formation)
	return &formation
}
func AfficherUserFor(user_id int) []*Formation {
	formation := []*Formation{}
	GetDB().Table("formations").Where(" user_id =?", user_id).Find(&formation)
	return formation
}

func ReserverFormation(participants []User, formationID int) *Formation {
	var formation Formation
	formation.Id = formationID

	if err := db.First(&formation).Error; err != nil {
		panic(err)
	}
	db.Model(&formation).Association("Participants").Append(participants)
	return &formation

}
