package main

import (
	_"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kotswane/tech_assessment_ioco_mtn/controller"
)

func main() {
	router := SetupRouter()
	router.Run(":8081")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
  
	v1 := router.Group("api/v1")
	{
	  v1.POST("/survivor", controller.RegisterSurvivor)
	  v1.POST("/survivor_flag", controller.FlagSurvivor)
	  v1.PUT("/survivor_location/:id", controller.UpdateLocationSurvivor)
	  v1.GET("/robot_list", controller.ListRobots)
	  v1.GET("/report_list_infected", controller.ListInfectedSurvivor)
	  v1.GET("/report_list_uninfected", controller.ListUnInfectedSurvivor)
	  v1.GET("/report_infected_percent", controller.GetInfectedPercent)
	  v1.GET("/report_uninfected_percent", controller.GetUnInfectedPercent)
	}
	return router
  }