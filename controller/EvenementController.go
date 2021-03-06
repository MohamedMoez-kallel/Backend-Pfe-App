package controller

import (
	"encoding/json"
	"net/http"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var AjouterEvenement = func(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")
	evenement := &model.Evenement{}
	err := json.NewDecoder(r.Body).Decode(evenement)
	if err != nil {
		u.Responds(w, u.Message("Requête invalide"))
		return
	}

	resp := evenement.AjouterEvenement(w, r)
	u.Respond(w, resp)
}
var AfficherEvenement = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherEvenement()
	resp := u.Message("Tous les evenements")
	resp["data"] = data

	u.Respond(w, resp)
}
var SupprimerEvenement = func(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.DeletEvenement(idInt)
	resp := u.Message("Evenement supprimer avec succès")
	resp["data"] = data
	u.Respond(w, resp)
}
var ModifierEvenement = func(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	evenement := &model.Evenement{}
	err1 := json.NewDecoder(r.Body).Decode(evenement)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	evenement.Id = idInt
	data := model.ModifierEvenement(evenement)
	resp := u.Message("Evenement modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}

var RechercheEvenement = func(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	title := vars["title"]
	evenement := model.RechercheEvenement(title)
	resp := u.Message("Evenement chercher ")
	resp["evenement"] = evenement
	u.Respond(w, resp)
}
var AfficherEvenementDate = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherEvenementDate()
	resp := u.Message("Evenement à partir du date d'aujourd'hui  ")
	resp["data"] = data

	u.Respond(w, resp)
}

var AfficherEve = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.AfficherParEve(idInt)
	resp := u.Message("Evenement")
	resp["data"] = data
	u.Respond(w, resp)
}
