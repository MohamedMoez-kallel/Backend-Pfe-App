package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"
)

var AddAnswer = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	answer := &model.Answer{}
	err1 := json.NewDecoder(r.Body).Decode(answer)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	data := answer.SaveComment(w, r)
	resp := u.Message("Commentaire ajoutée")
	resp["data"] = data

	u.Respond(w, resp)

}
var AfficherAnswer = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	pid, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.AfficheAnswer(pid)
	resp := u.Message("Tous les commentaire")
	resp["data"] = data
	u.Respond(w, resp)

}

// var AfficheAnswer = func(w http.ResponseWriter, r *http.Request) {

// 	data := model.AfficheAnswer()
// 	resp := u.Message("Tous les equipements")
// 	resp["data"] = data
// 	u.Respond(w, resp)
// }

var AfficherAns = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.AfficherParAns(idInt)
	resp := u.Message("Answer")
	resp["data"] = data
	u.Respond(w, resp)
}
var SupprimerAnswer = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.DeletAnswer(idInt)
	resp := u.Message("Commentaire supprimer avec succès")
	resp["data"] = data
	u.Respond(w, resp)
}

var ModifierAnswer = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	answer := &model.Answer{}
	err1 := json.NewDecoder(r.Body).Decode(answer)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	answer.Id = idInt
	data := model.ModifierAnswer(answer)
	resp := u.Message("Evenement modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}

type PublicationAnswer struct {
	PublicationID int            `json:"publication_id"`
	Reponse       []model.Answer `json:"response"`
}
