package model
import (
		"gorm.io/gorm"
)

type Survivors struct{
	gorm.Model
	ID 			int64 		 `json:"id"`
	Name 		string 		 `json:"name"`
	Age 		string 		 `json:"age"`
	Gender 		string       `json:"gender"`
	Latitude  	string       `json:"latitude"`
	Longitude   string 		 `json:"longitude"`
	Infected 	int64 		 `json:"infected"`
	Inventory  Inventories   `json:"inventory" gorm:"foreignKey:SurID"`
}