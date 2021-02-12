package controller

import (
	"encoding/json"
	"net/http"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var AjouterEquipement = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	equipement := &model.Equipement{}
	err := json.NewDecoder(r.Body).Decode(equipement)
	if err != nil {
		u.Responds(w, u.Messages("Requête invalide"))
		return
	}

	resp := equipement.AjouterEquipement(w, r)
	u.Respond(w, resp)
}

var AfficherEquipement = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherEquipement()
	resp := u.Message("Tous les equipements")
	resp["data"] = data
	u.Respond(w, resp)
}
var AfficherDisEquipement = func(w http.ResponseWriter, r *http.Request, nb_equipement int) {

	data := model.AfficherDisEquipement(nb_equipement)
	resp := u.Message("Tous les equipements disponible")
	resp["data"] = data
	u.Respond(w, resp)
}
var SupprimerEquipement = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.DeletEquipement(idInt)
	resp := u.Message("Equipement supprimer avec succès")
	resp["data"] = data
	u.Respond(w, resp)
}

var ModifierEquipement = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	equipement := &model.Equipement{}
	err1 := json.NewDecoder(r.Body).Decode(equipement)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	equipement.Id = idInt
	data := model.ModifierEquipement(equipement)
	resp := u.Message("Equipement modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}

var RechercheEquipement = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	title := vars["title"]
	equipement := model.RechercheEquipement(title)
	resp := u.Message("Equipement chercher")
	resp["equipement"] = equipement
	u.Respond(w, resp)
}
var AfficherEqui = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.AfficherParEqui(idInt)
	resp := u.Message("Equipement")
	resp["data"] = data
	u.Respond(w, resp)
}

var AfficherUserEqui = func(w http.ResponseWriter, r *http.Request) {

	paramss := mux.Vars(r)
	id_user := paramss["user_id"]
	idUser, err := strconv.Atoi(id_user)
	if err != nil {
	}
	data := model.AfficherUserEqui(idUser)
	resp := u.Message("Formations")
	resp["data"] = data
	u.Respond(w, resp)
}
