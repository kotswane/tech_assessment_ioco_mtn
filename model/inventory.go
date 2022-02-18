package model
import (
	"gorm.io/gorm"
)

type Inventories struct{
	gorm.Model
	ID 			int64  `json:"-" bson="id"`
	SurID 		string `json:"-" bson="sur_id"`
	Water 		string `json:"water"`
	Food 		string `json:"food"`
	Medication  string `json:"medication"`
	Ammunition  string `json:"ammunition"`
}