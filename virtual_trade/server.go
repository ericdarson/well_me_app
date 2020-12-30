package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"virtual_trade_api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/reksadana", handler.GetJenisReksadanaHandler())
	server.GET("/products/:id-jenis", handler.GetListProdukReksadanaHandler())
	server.GET("/product/:id-produk/:filter", handler.GetDetailProdukReksadanaHandler())
	server.GET("/simulation/start/:id-produk/:jumlahinvest", handler.SimulationStartHandler())
	server.GET("/simulation/forward/:id-produk/:simulation-date", handler.SimulationHandler())
	server.GET("/projection/:id-produk/:simulation-date", handler.ProjectionHandler())

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
