package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"rh-projet/model"
	u "rh-projet/utils"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

var AjouterUser = func(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		tel, _ := strconv.ParseInt(r.FormValue("tel"), 10, 64)
		salaire, _ := strconv.ParseFloat(r.FormValue("salaire"), 64)
		photo, err := UploadFileHandler(r)
		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user := &model.User{
			Nom:                r.FormValue("nom"),
			Prenom:             r.FormValue("prenom"),
			Cin:                r.FormValue("cin"),
			Date_naissance:     r.FormValue("date_naissance"),
			Sexe:               r.FormValue("sexe"),
			Email:              r.FormValue("email"),
			Tel:                tel,
			Lieu:               r.FormValue("lieu"),
			Departement:        r.FormValue("departement"),
			Type_contrat:       r.FormValue("type_contrat"),
			Date_debut_travail: r.FormValue("date_debut_travail"),
			Salaire:            salaire,
			Password:           r.FormValue("password"),
			Photo:              photo.FileName,
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
		resp := user.InsertUser(w, r)
		u.Respond(w, resp)
	}
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

func UploadFileHandler(r *http.Request) (*UploadedFile, error) {
	var acceptedExtensions = []string{".jpg", ".JPG", ".JPEG", ".jpeg", ".PNG", ".png"}
	if err := r.ParseMultipartForm(128 << 20); err != nil {
		return nil, err
	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	ext := path.Ext(header.Filename)
	if !contains(acceptedExtensions, ext) {
		return nil, errors.New("wrong image extension")
	}
	guid := xid.New()	
	filename := guid.String()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	uploadedFile := &UploadedFile{MultipartFile: file, FileName: filename, Extension: ext, FileBytes: fileBytes}
	return uploadedFile, nil
}

type UploadedFile struct {
	MultipartFile multipart.File
	FileName      string
	Extension     string
	FileBytes     []byte
}

func contains(s []string, ext string) bool {
	for _, e := range s {
		if strings.EqualFold(e, ext) {
			return true
		}
	}
	return false
}

func ErrorToJSON(err interface{}) ([]byte, error) {
	ErrorMap := make(map[string]interface{})
	ErrorMap["error"] = err
	status, er := json.Marshal(ErrorMap)
	if er != nil {
		return nil, er
	}
	return status, er
}
