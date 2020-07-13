package controller

import (
	"encoding/json"
	"net/http"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type SalleBody struct {
	Date_debut  string    `json:"date_debut"`
	Heure_debut time.Time `json:"heure_debut"`
	Heure_fin   time.Time `json:"heure_fin"`
	UserId      int       `json:"user_id"`
}

var AjouterSalle = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	salle := &model.Salle_reunion{}
	err := json.NewDecoder(r.Body).Decode(salle)
	if err != nil {
		u.Responds(w, u.Messages("Requête invalide"))
		return
	}

	resp := salle.InsertSalleReunion(w, r)
	u.Respond(w, resp)
}
var AfficherSalle = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherSalle()
	resp := u.Message("Tous les salle de réunions")
	resp["data"] = data
	u.Respond(w, resp)

}

var SupprimerSalle = func(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {

	}
	data := model.SupprimerSalle(idInt)
	resp := u.Message("Salle supprimer")
	resp["data"] = data
	u.Respond(w, resp)
}

var ModifierSalle = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	salle := &model.Salle_reunion{}
	err1 := json.NewDecoder(r.Body).Decode(salle)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	salle.Id = idInt
	data := model.ModifierSalle(salle)
	resp := u.Message("Salle de reunion modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}
var ReserverSalle = func(w http.ResponseWriter, r *http.Request) {

	//parking := &model.Parking{}
	var b *SalleBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		u.Responds(w, u.Messages("Requête invalide"))
		return
	}
	//fmt.Println(b.UserId)
	data := model.ReserverSalle(b.UserId, b.Date_debut, b.Heure_debut, b.Heure_fin)
	resp := u.Message("Salle réserver")
	resp[""] = data
	u.Respond(w, resp)
}
