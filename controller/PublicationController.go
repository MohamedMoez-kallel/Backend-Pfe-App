package controller

import (
	"encoding/json"
	"net/http"
	u "rh-projet/utils"
	"strconv"
	"rh-projet/model"

	
	"github.com/gorilla/mux"
)

var AjouterPublication = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	publication := &model.Publication{}
	err := json.NewDecoder(r.Body).Decode(publication)
	if err != nil {
		u.Responds(w, u.Messages("Requête invalide"))
		return
	}

	resp := publication.InsertPublication(w, r)
	u.Respond(w, resp)
}

var AfficherPublication = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherPublication()
	resp := u.Message("Tous les publications")
	resp["data"] = data
	u.Respond(w, resp)
}
var SupprimerPublication = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	//fmt.Println(idInt)
	data := model.DeletPublication(idInt)
	resp := u.Message("Publication supprimer avec succès")
	resp["data"] = data
	u.Respond(w, resp)
}
var AfficherPub = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.AfficherParPub(idInt)
	resp := u.Message("Publication")
	resp["data"] = data
	u.Respond(w, resp)
}
var AfficherUserPub = func(w http.ResponseWriter, r *http.Request) {

	paramss := mux.Vars(r)
	id_user := paramss["user_id"]
	idUser, err := strconv.Atoi(id_user)
	if err != nil {
	}
	data := model.AfficherUserPub(idUser)
	resp := u.Message("Publication")
	resp["data"] = data
	u.Respond(w, resp)
}

var ModifierPublications = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	publication := &model.Publication{}
	err1 := json.NewDecoder(r.Body).Decode(publication)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	publication.Id = idInt
	data := model.ModifierPublication(publication)
	resp := u.Message("Publication modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}
var RecherchePublication = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	title := vars["title"]
	publication := model.RecherchePublication(title)
	resp := u.Message("Publication chercher")
	resp["Publication"] = publication
	u.Respond(w, resp)
}
