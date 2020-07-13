package controller

import (
	"encoding/json"
	"net/http"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var AjouterUser = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Responds(w, u.Message("Requête invalide"))
		return
	}
	resp := user.InsertUser(w, r)
	u.Respond(w, resp)
}
var AfficherUser = func(w http.ResponseWriter, r *http.Request) {

	data := model.AfficherUser()
	resp := u.Message("Tous les utilisateurs")
	resp["data"] = data

	u.Respond(w, resp)
}

var SupprimerUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.DeleteUser(idInt)
	resp := u.Message("Utilisateur supprimer avec Succès")
	resp["data"] = data
	u.Respond(w, resp)
}
var ModifierUtilisateur = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	user := &model.User{}
	err1 := json.NewDecoder(r.Body).Decode(user)
	if err1 != nil {
		u.Responds(w, u.Messages("Requête invalide"))

		return
	}
	user.Id = idInt
	data := model.ModifierUtilisateur(user)
	resp := u.Message("Utilisateur modifier avec succès")
	resp["data"] = data
	u.Respond(w, resp)

}

var RechercheUtilisateur = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	user := model.RechercheUtilisateur(title)
	resp := u.Message("Utilisateur chercher")
	resp["user"] = user
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Responds(w, u.Message("Invalid request"))
		return
	}

	resp := model.Login(w, r, user.Email, user.Password, user.Nom, user.Departement)
	u.Respond(w, resp)
}
var AfficherUserById = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
	}
	data := model.AfficherParUser(idInt)
	resp := u.Message("User")
	resp["data"] = data
	u.Respond(w, resp)
}
