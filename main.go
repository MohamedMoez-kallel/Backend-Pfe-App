package main

import (
	"fmt"
	"net/http"
	"os"
	"rh-projet/controller"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	//USER//
	router.HandleFunc("/new/utilisateur", controller.AjouterUser).Methods("POST")
	router.HandleFunc("/afficher/utilisateur", controller.AfficherUser).Methods("GET")
	router.HandleFunc("/supprimer/utilisateur/{id}", controller.SupprimerUser).Methods("DELETE")
	router.HandleFunc("/modifier/utilisateur/{id}", controller.ModifierUtilisateur).Methods("PUT")
	router.HandleFunc("/chercher/utilisateur/{title}", controller.RechercheUtilisateur).Methods("GET")
	router.HandleFunc("/api/utilisateur/login", controller.Authenticate).Methods("POST")
	router.HandleFunc("/afficher/user/{id}", controller.AfficherUserById).Methods("GET")

	//PUBLICATION//
	router.HandleFunc("/new/publication", controller.AjouterPublication).Methods("POST")
	router.HandleFunc("/afficher/publication", controller.AfficherPublication).Methods("GET")
	router.HandleFunc("/supprimer/publication/{id}", controller.SupprimerPublication).Methods("DELETE")
	router.HandleFunc("/modifier/publication/{id}", controller.ModifierPublications).Methods("PUT")
	router.HandleFunc("/chercher/publication/{title}", controller.RecherchePublication).Methods("GET")
	router.HandleFunc("/afficher/pub/{id}", controller.AfficherPub).Methods("GET")

	//FORMATION//
	router.HandleFunc("/new/formation", controller.AjouterFormation).Methods("POST")
	router.HandleFunc("/afficher/formation", controller.AfficherFormation).Methods("GET")
	router.HandleFunc("/supprimer/formation/{id}", controller.SupprimerFormation).Methods("DELETE")
	router.HandleFunc("/modifier/formation/{id}", controller.ModifierFormation).Methods("PUT")
	router.HandleFunc("/chercher/formation/{title}", controller.RechercheFormation).Methods("GET")
	router.HandleFunc("/afficher/formation/date", controller.AfficherFormationDate).Methods("GET")
	router.HandleFunc("/afficher/formation/{id}", controller.AfficherFor).Methods("GET")

	//EVENEMENT//
	router.HandleFunc("/new/evenement", controller.AjouterEvenement).Methods("POST")
	router.HandleFunc("/afficher/evenement", controller.AfficherEvenement).Methods("GET")
	router.HandleFunc("/supprimer/evenement/{id}", controller.SupprimerEvenement).Methods("DELETE")
	router.HandleFunc("/modifier/evenement/{id}", controller.ModifierEvenement).Methods("PUT")
	router.HandleFunc("/chercher/evenement/{title}", controller.RechercheEvenement).Methods("GET")
	router.HandleFunc("/afficher/evenement/date", controller.AfficherEvenementDate).Methods("GET")
	router.HandleFunc("/afficher/evenement/{id}", controller.AfficherEve).Methods("GET")

	//EQUIPEMENT//
	router.HandleFunc("/new/equipement", controller.AjouterEquipement).Methods("POST")
	router.HandleFunc("/afficher/equipement", controller.AfficherEquipement).Methods("GET")
	router.HandleFunc("/supprimer/equipement/{id}", controller.SupprimerEquipement).Methods("DELETE")
	router.HandleFunc("/modifier/equipement/{id}", controller.ModifierEquipement).Methods("PUT")
	router.HandleFunc("/chercher/equipement/{title}", controller.RechercheEquipement).Methods("GET")
	router.HandleFunc("/afficher/equipement/{id}", controller.AfficherEqui).Methods("GET")


	//PARKING//
	router.HandleFunc("/new/place/parking", controller.AjouterParking).Methods("POST")
	router.HandleFunc("/afficher/place/parking", controller.AfficherParking).Methods("GET")
	router.HandleFunc("/nombre/place", controller.Count).Methods("GET")
	router.HandleFunc("/affecter/place/parking", controller.Affecter).Methods("POST")
	router.HandleFunc("/supprimer/parking/{id}", controller.SupprimerParking).Methods("DELETE")
	router.HandleFunc("/modifier/parking/{id}", controller.ModifierParking).Methods("PUT")

	//SALLE REUNION//
	router.HandleFunc("/new/salle_reunion", controller.AjouterSalle).Methods("POST")
	router.HandleFunc("/afficher/salle", controller.AfficherSalle).Methods("GET")
	router.HandleFunc("/supprimer/salle/{id}", controller.SupprimerSalle).Methods("DELETE")
	router.HandleFunc("/modifier/salle/{id}", controller.ModifierSalle).Methods("PUT")
	router.HandleFunc("/reserver/salle", controller.ReserverSalle).Methods("POST")

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
