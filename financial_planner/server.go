package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"widget_nab_service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST("/insertPlanner", handler.InsertPlanner())
	server.GET("/getListPlanner", handler.ListPlanner())
	server.PUT("/updatePlanner", handler.UpdatePlanner())
	server.DELETE("/deletePlanner", handler.DeletePlanner())
	server.POST("/simulasiPlanner", handler.TargetSimulationPlanner())
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	dat, err := ioutil.ReadFile(dir + "/properties/port.properties")
	if err != nil {
		fmt.Println(err)
	}
	port := string(dat)
	server.Run(":" + port)
}
