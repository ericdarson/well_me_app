package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"widget_api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/widget-nab", handler.NabWidgetGetByIds())
	server.GET("/widget-daily-profit", handler.DailyProfitWidgetGetByIds())
	server.GET("/widget-progress", handler.PlannerProgressWidgetGetByIds())

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
