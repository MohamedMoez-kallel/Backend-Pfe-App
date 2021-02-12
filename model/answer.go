package model

import (
	"net/http"
)

type Answer struct {
	Id            int
	Commentaire   string      `json:"commentaire"`
	UserId        int         `json:"user_id"`
	PublicationId int         `json:"publication_id"`
	User          User        `json:"user"`
	Publication   Publication `json:"publication"`
}

func (c *Answer) SaveComment(w http.ResponseWriter, r *http.Request) *Answer {
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &Answer{}
	}
	if c.Id != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.UserId).Take(&c.User).Error
		if err != nil {
			return &Answer{}
		}
	}
	return c
}

// func GetAnswers(pid int) (*[]Answer, error) {

// 	answer := []Answer{}
// 	err := db.Debug().Model(&Answer{}).Where("publication_id = ?", pid).Order("created_at desc").Find(&answer).Error
// 	if err != nil {
// 		return &[]Answer{}, err
// 	}
// 	if len(answer) > 0 {
// 		for i, _ := range answer {
// 			err := db.Debug().Model(&User{}).Where("id = ?", answer[i].UserId).Take(&answer[i].User).Error
// 			if err != nil {
// 				return &[]Answer{}, err
// 			}
// 		}
// 	}
// 	return &answer, err
// }

func AfficheAnswer(pid int) []*Answer {

	answer := []*Answer{}
	err := db.Debug().Model(&Answer{}).Where("publication_id = ?", pid).Find(&answer).Error
	if err != nil {
		return nil
	}
	return answer
}

func AfficherParAns(id int) *Answer {
	var answer Answer
	GetDB().Table("answers").Where("id = ?", id).Find(&answer)
	return &answer
}

func DeletAnswer(id int) *Answer {

	var answer Answer
	GetDB().First(&answer, id)
	db.Delete(&answer)
	return &answer
}

func ModifierAnswer(answer *Answer) *Answer {

	db.Save(&answer)
	return answer
}
