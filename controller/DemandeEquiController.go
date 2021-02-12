package controller

import (
	"encoding/json"
	"net/http"
	u "rh-projet/utils"
	"strconv"
	"rh-projet/model"

	"github.com/gorilla/mux"
)

var DemanderEquipement = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	demandeEquipement := &model.DemandeEquipement{}
	err := json.NewDecoder(r.Body).Decode(demandeEquipement)
	if err != nil {
		u.Responds(w, u.Messages("Requête invalide"))
		return
	}

	resp := demandeEquipement.DemanderEquipement(w, r)
	u.Respond(w, resp)
}

var AfficherEquiDemander = func(w http.ResponseWriter, r *http.Request) {

	paramss := mux.Vars(r)
	id_user := paramss["user_id"]
	idUser, err := strconv.Atoi(id_user)
	if err != nil {
	}
	data := model.AfficherEquiDemander(idUser)
	resp := u.Message("Mes Equipements demander")
	resp["data"] = data
	u.Respond(w, resp)
}

var AfficherAll = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherAll()
	resp := u.Message("Tous les equipeent demandés")
	resp["data"] = data

	u.Respond(w, resp)
}
