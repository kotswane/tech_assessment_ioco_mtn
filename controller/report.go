package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kotswane/tech_assessment_ioco_mtn/service"
)


func ListInfectedSurvivor(c *gin.Context) {

	resp,err := service.ListInfectedSurvivor()
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
		
}

func ListUnInfectedSurvivor(c *gin.Context) {

	resp,err := service.ListUnInfectedSurvivor()
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
		
}

func GetInfectedPercent(g *gin.Context){

	resp := service.GetInfectedPercent()
	g.JSON(http.StatusOK, gin.H{"data": resp})
}

func GetUnInfectedPercent(g *gin.Context){
	resp := service.GetUnInfectedPercent()
	g.JSON(http.StatusOK, gin.H{"data": resp})	
}