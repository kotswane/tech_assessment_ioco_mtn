package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kotswane/tech_assessment_ioco_mtn/model"
	"github.com/kotswane/tech_assessment_ioco_mtn/service"
	"fmt"
)


func RegisterSurvivor(c *gin.Context) {

	var survivor model.Survivors

	if errX := c.ShouldBindJSON(&survivor); errX != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errX.Error()})
		return
	}
	
	err := service.RegisterSurvivor(survivor)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered"})
		
}


func FlagSurvivor(c *gin.Context){
	
	var reports model.Reports
	if errX := c.ShouldBindJSON(&reports); errX != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errX.Error()})
		return
	}
	
    fmt.Println(reports.SurID)
	err := service.FlagSurvivor(reports.SurID)
	if err == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "survivor not infected"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "survivor infected"})
}

func UpdateLocationSurvivor(c *gin.Context){
	var survivor model.Survivors
	
	if errX := c.ShouldBindJSON(&survivor); errX != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errX.Error()})
		return
	}
	 
	id := c.Param("id")

	err := service.UpdateLocationSurvivor(survivor,id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location updated"})
}