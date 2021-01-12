package main

import (
	"annual_report/handler"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/annual-report", handler.AnnualReportGetByBcaId())

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
