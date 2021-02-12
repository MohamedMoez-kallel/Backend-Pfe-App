package model

import (
	"net/http"
	u "rh-projet/utils"
)

type DemandeEquipement struct {
	Id                int
	Nb_equiDemander   int    `json:"nb_equi_demander"`
	Type_equiDemander string `json:"type_equi_demander"`
	Photo             string `json:"photo"`
	UserId            int    `json:"user_id"`
	User_nom          string `json:"user_nom"`
}

func (demanderEquipement *DemandeEquipement) DemanderEquipement(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	GetDB().Create(demanderEquipement)
	response := u.Message("Equipement demander avec succès")
	return response
}

func AfficherEquiDemander(user_id int) []*DemandeEquipement {
	demanderEquipement := []*DemandeEquipement{}
	GetDB().Table("demande_equipements").Where(" user_id =?", user_id).Find(&demanderEquipement)
	return demanderEquipement
}

func AfficherAll() []*DemandeEquipement {

	demanderEquipement := []*DemandeEquipement{}
	err := GetDB().Table("demande_equipements").Find(&demanderEquipement).Error
	if err != nil {
		return nil
	}

	return demanderEquipement
}
