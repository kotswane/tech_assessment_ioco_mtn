package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kotswane/tech_assessment_ioco_mtn/service"
	"github.com/kotswane/tech_assessment_ioco_mtn/utils"
	"strings"
	_"fmt"
)


func ListRobots(g *gin.Context){

	orderby := g.Query("orderby")
	orderist := []string{"model","serialNumber","manufacturedDate","category"}
	
	if orderby != "" {
		if(!utils.Contains(orderist,orderby)){
			g.JSON(http.StatusBadRequest, gin.H{"error": "Valid orderby [model, serialNumber, manufacturedDate, category]"})
			return
		}
	}	
	resp,err := service.ListRobots(strings.Title(orderby))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	g.JSON(http.StatusOK, gin.H{"data": resp})
}