package model

import (
	"net/http"
	u "rh-projet/utils"
)

type Parking struct {
	Id         int    `json:"id"`
	Matricule  string `json:"matricule"`
	Date_debut string `json:"date_debut"`
	Date_fin   string `json:"date_fin"`
	User       User
	UserId     int `json:"user_id"`
}

func (parking *Parking) InsertParking(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	GetDB().Create(parking)
	response := u.Message(" Place parking ajouter")
	return response
}

func AfficherParking() []*Parking {

	Parking := []*Parking{}
	err := GetDB().Table("parkings").Find(&Parking).Error
	if err != nil {
		return nil
	}

	return Parking
}

func (parking *Parking) Count() int64 {

	var result int64
	db.Table("parkings").Where("user_id = 0").Count(&result)
	return result
}

func AffecterPlaceParking(user_id int, Date_debut, Date_fin, Matricule string) *Parking {

	var result int
	parking := []*Parking{}
	db.Table("parkings").Where("user_id = 0").Count(&result)
	//fmt.Println(result)
	err := GetDB().Table("parkings").Where("user_id = 0").Find(&parking).Error
	if err != nil {
		return nil
	}
	if result > 0 {

		parking[0].UserId = user_id
		parking[0].Date_debut = Date_debut
		parking[0].Date_fin = Date_fin
		parking[0].Matricule = Matricule
		db.Table("parkings").Save(&parking[0])

	}
	return parking[0]
}
func SupprimerParking(id int) *Parking {
	var parking Parking
	GetDB().First(&parking, id)
	db.Delete(&parking)
	return &parking

}

func ModifierParking(parking *Parking) *Parking {

	db.Save(&parking)
	return parking
}

func AfficherParPlace(id int) *Parking {
	var parking Parking
	GetDB().Table("parkings").Where("id = ?", id).Find(&parking)
	return &parking
}
