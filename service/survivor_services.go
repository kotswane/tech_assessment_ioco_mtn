package service

import (
	"github.com/kotswane/tech_assessment_ioco_mtn/model"
	"github.com/kotswane/tech_assessment_ioco_mtn/config/db"
)

func RegisterSurvivor(survivor model.Survivors) error{

	var db = db.Init()
    result := db.Create(&survivor)

	if result.Error != nil{
		return result.Error
	}
	return nil
}

func FlagSurvivor(id string) bool{

	var count int64
	var reports model.Reports
	var db = db.Init()

	survivor := getById(id) 

	if survivor.ID == 0{
		return false
	}

	db.Model(&reports).Where("sur_id",id).Count(&count)
	if(count <= 2 ){
		reports.SurID = id
		err := db.Create(&reports)
		if err != nil {
			return false
		}
		return false

	}else{
		if(count >= 3){
			survivor.Infected = 1
			_ = db.Save(&survivor)
			return true
		}
	}	
	return false	
}

func UpdateLocationSurvivor(survivor model.Survivors, id string) error{
	
	oldSurvivor := getById(id)
	oldSurvivor.Latitude = survivor.Latitude
	oldSurvivor.Longitude = survivor.Longitude
	
	var db = db.Init()
    result := db.Save(&oldSurvivor)

	if result.Error != nil{
		return result.Error
	}
	return nil
}

func getById(id string) model.Survivors {
	var db = db.Init()
	var survivor model.Survivors
	db.First(&survivor, id)
	return survivor
}