package main

import (
	"fmt"
	"io/ioutil"
	"list_reksadana_api/handler"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/products/:id-jenis", handler.ListProdukReksadanaHandler())

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
