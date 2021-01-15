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

	server.GET("/widget-nab", handler.NabWidgetGetByIds())
	server.POST("/login", handler.Login())
	server.POST("/checkSession", handler.CheckSession())
	server.POST("/logout", handler.Logout())
	server.GET("/getProfile", handler.Profile())
	server.POST("/insertPlanner", handler.InsertPlanner())
	server.GET("/getListPlanner", handler.ListPlanner())
	server.PUT("/updatePlanner", handler.UpdatePlanner())
	server.DELETE("/deletePlanner", handler.DeletePlanner())
	server.POST("/simulasiPlanner", handler.TargetSimulationPlanner())
	server.POST("/loginAdmin", handler.LoginAdmin())
	server.GET("/checkSessionAdmin", handler.CheckSessionAdmin())
	server.GET("/logoutAdmin", handler.LogoutAdmin())
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
