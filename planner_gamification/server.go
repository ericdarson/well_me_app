package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"planner_gamification_api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/financial-planner/details/:id-plan", handler.PlannerGamificationHandler())

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
