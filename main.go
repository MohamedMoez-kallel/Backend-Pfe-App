package main

import (
	"fmt"
	"net/http"
	"os"
	"rh-projet/controller"

	// calendar "google.golang.org/api/calendar/v3"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	//USER//

	router.HandleFunc("/file", controller.UploadFiles).Methods("POST")
	router.HandleFunc("/UploadFiles", controller.UploadFiles).Methods("POST")

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
	router.HandleFunc("/afficher/Userpub/{user_id}", controller.AfficherUserPub).Methods("GET")

	//FORMATION//
	router.HandleFunc("/new/formation", controller.AjouterFormation).Methods("POST")
	router.HandleFunc("/afficher/formation", controller.AfficherFormation).Methods("GET")
	router.HandleFunc("/supprimer/formation/{id}", controller.SupprimerFormation).Methods("DELETE")
	router.HandleFunc("/modifier/formation/{id}", controller.ModifierFormation).Methods("PUT")
	router.HandleFunc("/chercher/formation/{title}", controller.RechercheFormation).Methods("GET")
	router.HandleFunc("/afficher/formation/date", controller.AfficherFormationDate).Methods("GET")
	router.HandleFunc("/afficher/formation/{id}", controller.AfficherFor).Methods("GET")
	router.HandleFunc("/afficher/Userfor/{user_id}", controller.AfficherUserFor).Methods("GET")
	router.HandleFunc("/reserver/place/formation", controller.Reserver).Methods("POST")

	//EVENEMENT//
	router.HandleFunc("/new/evenement", controller.AjouterEvenement).Methods("POST")
	router.HandleFunc("/afficher/evenement", controller.AfficherEvenement).Methods("GET")
	router.HandleFunc("/supprimer/evenement/{id}", controller.SupprimerEvenement).Methods("DELETE")
	router.HandleFunc("/modifier/evenement/{id}", controller.ModifierEvenement).Methods("PUT")
	router.HandleFunc("/chercher/evenement/{title}", controller.RechercheEvenement).Methods("GET")
	router.HandleFunc("/afficher/evenement/date", controller.AfficherEvenementDate).Methods("GET")
	router.HandleFunc("/afficher/evenement/{id}", controller.AfficherEve).Methods("GET")
	router.HandleFunc("/afficher/userEve/{user_id}", controller.AfficherUserEve).Methods("GET")
	router.HandleFunc("/reserver/place/evenement", controller.ReserverEvenement).Methods("POST")

	//EQUIPEMENT//
	router.HandleFunc("/new/equipement", controller.AjouterEquipement).Methods("POST")
	router.HandleFunc("/afficher/equipement", controller.AfficherEquipement).Methods("GET")
	router.HandleFunc("/supprimer/equipement/{id}", controller.SupprimerEquipement).Methods("DELETE")
	router.HandleFunc("/modifier/equipement/{id}", controller.ModifierEquipement).Methods("PUT")
	router.HandleFunc("/chercher/equipement/{title}", controller.RechercheEquipement).Methods("GET")
	router.HandleFunc("/afficher/equipement/{id}", controller.AfficherEqui).Methods("GET")
	router.HandleFunc("/afficher/equipements/{user_id}", controller.AfficherUserEqui).Methods("GET")

	//PARKING//
	router.HandleFunc("/new/place/parking", controller.AjouterParking).Methods("POST")
	router.HandleFunc("/afficher/place/parking", controller.AfficherParking).Methods("GET")
	router.HandleFunc("/nombre/place", controller.Count).Methods("GET")
	router.HandleFunc("/affecter/place/parking", controller.Affecter).Methods("POST")
	router.HandleFunc("/supprimer/parking/{id}", controller.SupprimerParking).Methods("DELETE")
	router.HandleFunc("/modifier/parking/{id}", controller.ModifierParking).Methods("PUT")
	router.HandleFunc("/afficher/place/{id}", controller.AfficherParPlace).Methods("GET")

	//SALLE REUNION//
	router.HandleFunc("/new/salle_reunion", controller.AjouterSalle).Methods("POST")
	router.HandleFunc("/afficher/salle", controller.AfficherSalle).Methods("GET")
	router.HandleFunc("/supprimer/salle/{id}", controller.SupprimerSalle).Methods("DELETE")
	router.HandleFunc("/modifier/salle/{id}", controller.ModifierSalle).Methods("PUT")
	router.HandleFunc("/reserver/salle", controller.ReserverSalle).Methods("POST")

	//Equipement Demander
	router.HandleFunc("/demander/equipement", controller.DemanderEquipement).Methods("POST")
	router.HandleFunc("/mes_equipements/{user_id}", controller.AfficherEquiDemander).Methods("GET")
	router.HandleFunc("/afficher/equi/demander", controller.AfficherAll).Methods("GET")

	//Answer

	router.HandleFunc("/add/answer/{id}", controller.AddAnswer).Methods("POST")
	router.HandleFunc("/afficher/answer/{id}", controller.AfficherAnswer).Methods("GET")
	// router.HandleFunc("/afficher/answer/{id}", controller.AfficherAns).Methods("GET")
	// router.HandleFunc("/supprimer/answer/{id}", controller.SupprimerAnswer).Methods("DELETE")
	// router.HandleFunc("/modifier/answer/{id}", controller.ModifierAnswer).Methods("PUT")

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

// func calendarMain(client *http.Client, argv []string) {
// 	if len(argv) != 0 {
// 		fmt.Fprintln(os.Stderr, "Usage: calendar")
// 		return
// 	}

// 	svc, err := calendar.New(client)
// 	if err != nil {
// 		log.Fatalf("Unable to create Calendar service: %v", err)
// 	}

// 	c, err := svc.Colors.Get().Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve calendar colors: %v", err)
// 	}

// 	log.Printf("Kind of colors: %v", c.Kind)
// 	log.Printf("Colors last updated: %v", c.Updated)

// 	for k, v := range c.Calendar {
// 		log.Printf("Calendar[%v]: Background=%v, Foreground=%v", k, v.Background, v.Foreground)
// 	}

// 	for k, v := range c.Event {
// 		log.Printf("Event[%v]: Background=%v, Foreground=%v", k, v.Background, v.Foreground)
// 	}

// 	listRes, err := svc.CalendarList.List().Fields("items/id").Do()
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve list of calendars: %v", err)
// 	}
// 	for _, v := range listRes.Items {
// 		log.Printf("Calendar ID: %v\n", v.Id)
// 	}

// 	if len(listRes.Items) > 0 {
// 		id := listRes.Items[0].Id
// 		res, err := svc.Events.List(id).Fields("items(updated,summary)", "summary", "nextPageToken").Do()
// 		if err != nil {
// 			log.Fatalf("Unable to retrieve calendar events list: %v", err)
// 		}
// 		for _, v := range res.Items {
// 			log.Printf("Calendar ID %q event: %v: %q\n", id, v.Updated, v.Summary)
// 		}
// 		log.Printf("Calendar ID %q Summary: %v\n", id, res.Summary)
// 		log.Printf("Calendar ID %q next page token: %v\n", id, res.NextPageToken)
// 	}
// }
