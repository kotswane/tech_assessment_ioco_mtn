package service

import (
	"github.com/kotswane/tech_assessment_ioco_mtn/model"
	"github.com/kotswane/tech_assessment_ioco_mtn/config/db"
)

func ListInfectedSurvivor() ([]model.Survivors, error){
	var db = db.Init()
	var survivor []model.Survivors
	db.Where("infected",1).Find(&survivor)
	return survivor,nil
}

func ListUnInfectedSurvivor() ([]model.Survivors, error){
	var db = db.Init()
	var survivor []model.Survivors
	db.Where("infected",0).Find(&survivor)
	return survivor,nil
}

func GetInfectedPercent() int64{
	var db = db.Init()
	var survivor []model.Survivors

	var countInfected int64
	var totalRecords int64
	var percent int64

	db.Model(&survivor).Where("infected",1).Count(&countInfected)

	if countInfected == 0{
		return 0
	}
	db.Table("Survivors").Count(&totalRecords)
   
	percent = ((countInfected * 100)/totalRecords)
	return percent
}

func GetUnInfectedPercent() int64{

	var db = db.Init()
	var survivor []model.Survivors

	var countInfected int64
	var totalRecords int64
	var percent int64

	db.Model(&survivor).Where("infected",0).Count(&countInfected)
	db.Table("Survivors").Count(&totalRecords)
	if countInfected == 0{
		return 0
	}
	percent = ((countInfected * 100)/totalRecords)
	return percent
}