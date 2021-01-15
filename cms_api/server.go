package main

import (
	"cms_api/handler"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//promo
	server.GET("/promo", handler.InquiryPromo())
	server.POST("/promo/akumulasi/", handler.InsertPromoAkumulasi())
	server.POST("/promo/transaksi/", handler.InsertPromoTransaksi())
	server.DELETE("/promo/:kode-promo", handler.NonAktifPromo())
	server.PUT("/promo/akumulasi/:kode-promo", handler.UpdatePromoAkumulasi())
	server.PUT("/promo/transaksi/:kode-promo", handler.UpdatePromoTransaksi())

	//jenis reksadana
	server.GET("/jenis-reksadana", handler.InquiryAllJenisReksadana())
	server.POST("/jenis-reksadana", handler.InsertJenisReksadana())
	server.PUT("/jenis-reksadana/:id-jenis", handler.UpdateJenisReksadana())

	//product reksadana
	server.GET("/produk-reksadana", handler.InquiryProdukReksadana())
	server.POST("/produk-reksadana", handler.InsertProdukReksadana())
	server.PUT("/produk-reksadana/:id-produk", handler.UpdateProdukReksadana())

	//resiko
	server.GET("/bobot-resiko", handler.InquiryBobotResiko())
	server.PUT("/bobot-resiko", handler.UpdateBobotResiko())

	//dashboard
	server.GET("/dashboard/overview/:chart-type/:start-date/:end-date", handler.InquiryDashboardOverview())
	server.GET("/dashboard/promo", handler.InquiryPromoUsage())

	//history transaksi
	server.GET("/transaction-history/pembelian/:bca-id", handler.PurchaseHistory())
	server.GET("/transaction-history/penjualan/:bca-id", handler.SellHistory())
	server.GET("/transaction-history/pembelian", handler.PurchaseHistory())
	server.GET("/transaction-history/penjualan", handler.SellHistory())

	//nab
	server.GET("/daily-nab/products/:id-jenis", handler.GetDailyProducts())
	server.POST("/daily-nab/products", handler.InsertDailyProducts())

	//nasabah
	server.GET("/nasabah/:bca-id", handler.GetDataNasabah())

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
