package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var AjouterFormation = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	formation := &model.Formation{}
	err := json.NewDecoder(r.Body).Decode(formation)
	if err != nil {
		u.Responds(w, u.Messages("Requête invalide"))
		return
	}

	resp := formation.InsertFormation(w, r)
	u.Respond(w, resp)
}

var AfficherFormation = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherFormation()
	resp := u.Message("Tous les formations")
	resp["data"] = data
	u.Respond(w, resp)

}
var SupprimerFormation = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	fmt.Println(idInt)
	data := model.SupprimerFormation(idInt)
	resp := u.Message("Formation supprimer avec Succès")
	resp["data"] = data
	u.Respond(w, resp)
}
var ModifierFormation = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	formation := &model.Formation{}
	err1 := json.NewDecoder(r.Body).Decode(formation)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	formation.Id = idInt
	data := model.ModifierFormation(formation)
	resp := u.Message("Formation modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}
var RechercheFormation = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	title := vars["title"]
	formation := model.RechercheFormation(title)
	resp := u.Message("Formation chercher")
	resp["Formation"] = formation
	u.Respond(w, resp)
}
var AfficherFormationDate = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherFormationDate()
	resp := u.Message("Formation à partir du date d'aujourd'hui  ")
	resp["data"] = data

	u.Respond(w, resp)
}

var AfficherFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.AfficherParFor(idInt)
	resp := u.Message("Formation")
	resp["data"] = data
	u.Respond(w, resp)
}
