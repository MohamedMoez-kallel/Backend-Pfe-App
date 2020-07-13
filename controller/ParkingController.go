package controller

import (
	"encoding/json"
	"net/http"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var AjouterParking = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	parking := &model.Parking{}
	err := json.NewDecoder(r.Body).Decode(parking)
	if err != nil {
		u.Responds(w, u.Message("Requête invalide"))
		return
	}

	resp := parking.InsertParking(w, r)
	u.Respond(w, resp)
}
var AfficherParking = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherParking()
	resp := u.Message("Tous les place parking")
	resp["data"] = data

	u.Respond(w, resp)
}
var Count = func(w http.ResponseWriter, r *http.Request) {

	parking := &model.Parking{}
	data := parking.Count()
	resp := u.Message("")
	resp["Nombre de place est ="] = data
	u.Respond(w, resp)
}

type ParkingBody struct {
	Matricule  string `json:"matricule"`
	Date_debut string `json:"date_debut"`
	Date_fin   string `json:"date_fin"`
	UserId     int    `json:"user_id"`
}

var Affecter = func(w http.ResponseWriter, r *http.Request) {

	//parking := &model.Parking{}
	var b *ParkingBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		u.Responds(w, u.Messages("Requête invalide"))
		return
	}
	//fmt.Println(b.UserId)
	data := model.AffecterPlaceParking(b.UserId,b.Date_debut,b.Date_fin,b.Matricule)
	resp := u.Message("Utilisateur affecter à une place parking")
	resp[""] = data
	u.Respond(w, resp)
}
var SupprimerParking = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.SupprimerParking(idInt)
	resp := u.Message("Place parking supprimer avec Succès")
	resp["data"] = data
	u.Respond(w, resp)
}
var ModifierParking = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	parking := &model.Parking{}
	err1 := json.NewDecoder(r.Body).Decode(parking)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	parking.Id = idInt
	data := model.ModifierParking(parking)
	resp := u.Message("Utilisateur modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}
