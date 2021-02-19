package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"
)

var AjouterFormation = func(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		nb_place, _ := strconv.ParseInt(r.FormValue("Nb_place"), 10, 64)
		photo, err := UploadFileHandler(r)
		if err != nil {
			return
		}
		formation := &model.Formation{
			Heure_formation: r.FormValue("heure_formation"),
			Date_formation:  r.FormValue("date_formation"),
			Titre:           r.FormValue("titre"),
			Description:     r.FormValue("description"),
			Formateur:       r.FormValue("formateur"),
			Nb_place:        nb_place,
			Type_formation:  r.FormValue("type_formation"),
			Photo:           photo.FileName,
		}
		newPath := filepath.Join(".", "C:/Users/Moez/Rh-Projet-Client/src/assets/image/")
		err = os.MkdirAll(newPath, os.ModePerm)
		if err != nil {
			fmt.Println((err))
		}
		err = ioutil.WriteFile(newPath+"/"+photo.FileName+photo.Extension, photo.FileBytes, 0644)
		if err != nil {
			fmt.Println((err))
		}
		resp := formation.InsertFormation(w, r)
		u.Respond(w, resp)
	}
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
	resp := u.Message("Formation à partir du date d'aujourd'hui")
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
var AfficherUserFor = func(w http.ResponseWriter, r *http.Request) {

	paramss := mux.Vars(r)
	id_user := paramss["user_id"]
	idUser, err := strconv.Atoi(id_user)
	if err != nil {
	}
	data := model.AfficherUserFor(idUser)
	resp := u.Message("Formations")
	resp["data"] = data
	u.Respond(w, resp)
}

type ParticipationRequest struct {
	FormationID  int          `json:"formation_id"`
	Participants []model.User `json:"participants"`
}

var Reserver = func(w http.ResponseWriter, r *http.Request) {

	var b *ParticipationRequest
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		u.Responds(w, u.Messages("Requête invalide"))
		return
	}
	//fmt.Println(b.UserId)
	data := model.ReserverFormation(b.Participants, b.FormationID)
	resp := u.Message("Utilisateur a reserver une place")
	resp["data"] = data
	u.Respond(w, resp)
}
