package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pembelian_verifikasi_bank_api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.PUT("/pembelian/verifikasi-bank/:id-transaksi", handler.PlannerVerifikasiBankPembelianHandler())

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
