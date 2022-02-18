package service

import (
		"net/http"
		"io"
		"github.com/kotswane/tech_assessment_ioco_mtn/model"
		"github.com/kotswane/tech_assessment_ioco_mtn/config/httpconfig"
		"encoding/json"
		"sort"
	)

func ListRobots(orderby string) ([]model.Robot,error){

	conf := httpconfig.Init()
	var client http.Client
	var robot []model.Robot
	
	resp, err := client.Get(conf)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil,err
		}

		json.Unmarshal([]byte(bodyBytes), &robot)
		if(orderby == "Category"){
			sort.Slice(robot, func(i, j int) bool {
				return robot[i].Category < robot[j].Category
			})
		}else if(orderby == "Model"){
			sort.Slice(robot, func(i, j int) bool {
				return robot[i].Model < robot[j].Model
			})
		}else if(orderby == "SerialNumber"){
			sort.Slice(robot, func(i, j int) bool {
				return robot[i].SerialNumber < robot[j].SerialNumber
			})
		}else if(orderby == "ManufacturedDate"){
			sort.Slice(robot, func(i, j int) bool {
				return robot[i].ManufacturedDate < robot[j].ManufacturedDate
			})
		}
		return robot,nil
	}
	return nil,nil
}