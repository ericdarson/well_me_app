package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pembelian_finish_api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.PUT("/pembelian/finish/:id-transaksi", handler.PembelianFinish())

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
