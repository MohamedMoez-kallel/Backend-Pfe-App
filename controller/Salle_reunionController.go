package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ReservationBody struct {
	Date_debut  string `json:"date_debut"`
	Heure_debut string `json:"heure_debut"`
	Heure_fin   string `json:"heure_fin"`
	UserId      int    `json:"user_id"`
	SalleId     int    `json:"salle_id"`
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
	salle.ID = uint(idInt)
	data := model.ModifierSalle(salle)
	resp := u.Message("Salle de reunion modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}
var ReserverSalle = func(w http.ResponseWriter, r *http.Request) {
	var b *ReservationBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		u.Response(w, http.StatusBadRequest, errors.New("Requête invalide").Error())
		return
	}
	heureDebut, err := time.Parse(time.RFC3339, b.Heure_debut)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed heure debut %v", err))
		u.Response(w, http.StatusInternalServerError, err.Error())
		return
	}
	heureFin, err := time.Parse(time.RFC3339, b.Heure_fin)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed heure fin %v", err))
		u.Response(w, http.StatusInternalServerError, err.Error())
		return
	}
	if heureDebut.After(heureFin) {
		u.Response(w, http.StatusInternalServerError, errors.New("wrong time slot").Error())
		return
	}
	data, err := model.ReserverSalle(b.SalleId, b.UserId, b.Date_debut, heureDebut, heureFin)
	if err != nil {
		u.Response(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.Response(w, http.StatusOK, data)
}
