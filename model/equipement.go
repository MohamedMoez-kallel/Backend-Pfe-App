package model

import (
	"net/http"
	u "rh-projet/utils"
)

type Equipement struct {
	//gorm.Model

	Id              int
	Nb_equipement   int     `json:"nb_equipement"`
	Type_equipement string  `json:"type_equipement"`
	Photo           string  `json:"photo"`
	Prix            float64 `json:"prix"`
	Fournisseur     string  `json:"fournisseur"`
	Num_serie       string  `json:"num_serie"`
}

func (equipement *Equipement) AjouterEquipement(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	GetDB().Create(equipement)
	response := u.Message("Equipement ajouter avec succès")
	return response
}

func AfficherEquipement() []*Equipement {

	equipement := []*Equipement{}
	err := GetDB().Table("equipements").Find(&equipement).Error
	if err != nil {
		return nil
	}
	return equipement
}

func DeletEquipement(id int) *Equipement {

	var equipement Equipement
	GetDB().First(&equipement, id)
	db.Delete(&equipement)
	return &equipement
}

func ModifierEquipement(equipement *Equipement) *Equipement {

	db.Save(&equipement)
	return equipement
}

func RechercheEquipement(title string) []*Equipement {

	equipement := []*Equipement{}
	GetDB().Table("equipements").Where("type_equipement LIKE ? or num_serie LIKE ? or fournisseur LIKE ? ", title, title, title).Find(&equipement)

	return equipement
}
func AfficherParEqui(id int) *Equipement {
	var equipement Equipement
	GetDB().Table("equipements").Where("id = ?", id).Find(&equipement)
	return &equipement
}
