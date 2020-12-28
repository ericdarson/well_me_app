package main

import (
	"widget_nab_service/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/widget-nab", handler.NabWidgetGetByIds())

	server.Run(":8080")
}
