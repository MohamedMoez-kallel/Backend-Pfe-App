package model

import (
	"net/http"
	"os"
	u "rh-projet/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId int
	jwt.StandardClaims
}

type User struct {
	//gorm.Model
	Id                 int     `json:"id"`
	Nom                string  `json:"nom"`
	Prenom             string  `json:"prenom"`
	Date_naissance     string  `json:"date_naissance"`
	Cin                string  `json:"cin"`
	Departement        string  `json:"departement"`
	Type_contrat       string  `json:"type_contrat"`
	Salaire            float64 `json:"salaire"`
	Date_debut_travail string  `json:"date_debut_travail"`
	Email              string  `json:"email"`
	Photo              string  `json:"photo"`
	Tel                int64   `json:"tel"`
	Transport          string  `json:"transport"`
	Password           string  `json:"password"`
	Token              string  `json:"token"`
	Sexe               string  `json:"sexe"`
	Lieu               string  `json:"lieu"`
	Avantage           string  `json:"avantage"`
}

func (user *User) Validate(w http.ResponseWriter, r *http.Request) (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		w.WriteHeader(http.StatusBadRequest)
		return u.Messages("Adresse e-mail invalide"), false
	}
	if len(user.Password) < 6 {
		w.WriteHeader(http.StatusBadRequest)
		return u.Messages("Mot de passe invalide"), false
	}
	test := &User{}
	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(test).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusBadRequest)
		return u.Message("Connection error. Please retry"), false
	}
	if test.Email != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")

		return u.Messages("Adresse e-mail est utilisée par un autre utilisateur"), false
	}
	return u.Messages("Requirement passed"), true
}

func (user *User) InsertUser(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	if resp, ok := user.Validate(w, r); !ok {
		return resp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	
	GetDB().Create(user)
	user.Password = "" //delete password
	response := u.Message("Compte a été créé")
	return response
}

func AfficherUser() []*User {

	user := []*User{}
	err := GetDB().Table("users").Find(&user).Error
	if err != nil {
		return nil
	}
	return user
}

func DeleteUser(id int) *User {

	var user User
	GetDB().First(&user, id)
	db.Delete(&user)
	return &user
}

func ModifierUtilisateur(user *User) *User {

	db.Save(&user)
	return user
}

func RechercheUtilisateur(title string) []*User {

	user := []*User{}
	GetDB().Table("users").Where("nom LIKE ? or prenom LIKE ? or departement LIKE ? or type_contrat LIKE ? ", title, title, title, title).Find(&user)
	return user
}

func Login(w http.ResponseWriter, r *http.Request, email, password, nom, departement string) map[string]interface{} {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//w.WriteHeader(http.StatusBadRequest)
			w.WriteHeader(http.StatusNotFound)
			w.Header().Add("Content-Type", "application/json")
			return u.Messages("Adresse e-mail introuvable")
		}
		return u.Message("Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")
		return u.Messages("Invalid login credentials. Please try again")
	}

	user.Password = ""
	expirationTime := time.Now().Add(60 * time.Minute)
	tk := &Token{UserId: user.Id, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response
	resp := u.Message("Logged In")
	resp["Id"] = user.Id
	resp["Welcome"] = user.Nom
	resp["Departement"] = user.Departement
	resp["Email"] = user.Email
	resp["Token"] = user.Token
	resp["Nom"] = user.Nom

	return resp
}
func AfficherParUser(id int) *User {
	var user User
	GetDB().Table("users").Where("id = ?", id).Find(&user)
	return &user
}
