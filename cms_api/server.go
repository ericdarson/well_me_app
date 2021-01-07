package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"promo_cms/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/promo", handler.InquiryPromo())
	server.POST("/promo/akumulasi/", handler.InsertPromoAkumulasi())
	server.POST("/promo/transaksi/", handler.InsertPromoTransaksi())
	server.DELETE("/promo/:kode-promo", handler.NonAktifPromo())
	server.PUT("/promo/akumulasi/:kode-promo", handler.UpdatePromoAkumulasi())
	server.PUT("/promo/transaksi/:kode-promo", handler.UpdatePromoTransaksi())

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
