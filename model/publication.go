package model

import (
	"fmt"
	"net/http"
	u "rh-projet/utils"
)

type Publication struct {
	//gorm.Model
	Id           int
	Titre        string `json:"titre"`
	Description  string `json:"description"`
	NumberOfLike int    `json:"numberOfLike"`
	User         User
	UserId       int `json:"user_id"`
	PublicationAnswer []Answer `gorm:"many2many:publication_answer;"`

}

func (publication *Publication) InsertPublication(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	GetDB().Create(publication)
	response := u.Message("Publication ajouter avec succès")
	return response
}

func AfficherPublication() []*Publication {

	publication := []*Publication{}
	err := GetDB().Table("publications").Find(&publication).Error
	if err != nil {
		return nil
	}
	return publication
}

func AfficherPub(id int) []*Publication {
	publications := make([]*Publication, 0)
	err := GetDB().Table("publications").Where("id = ?", id).Find(&publications).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return publications
}

func AfficherParPub(id int) *Publication {
	var publication Publication
	GetDB().Table("publications").Where("id = ?", id).Find(&publication)
	return &publication
}

func AfficherUserPub(user_id int) []*Publication {
	publication := []*Publication{}
	GetDB().Table("publications").Where(" user_id =?", user_id).Find(&publication)
	return publication
}

func DeletPublication(id int) *Publication {

	var publication Publication
	GetDB().First(&publication, id)
	db.Delete(&publication)
	return &publication
}

func ModifierPublication(publication *Publication) *Publication {

	db.Save(&publication)
	return publication
}

func RecherchePublication(title string) []*Publication {

	publication := []*Publication{}
	GetDB().Table("publications").Where("titre LIKE ? ", title).Find(&publication)
	return publication
}

