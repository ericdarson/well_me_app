package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"promo_activate/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.POST("/promo/activate/", handler.ActivatePromo())
	server.GET("/promo/:bca-id", handler.InquiryPromo())
	server.PUT("/promo/claim/", handler.ClaimPromo())

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
