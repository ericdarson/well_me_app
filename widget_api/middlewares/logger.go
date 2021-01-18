package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerToFile() gin.HandlerFunc {

	//write file
	src, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("err", err)
	}
	//instantiation
	logger := logrus.New()
	//Set output
	logger.Out = src
	//Set log level
	logger.SetLevel(logrus.DebugLevel)
	//Format log
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// End time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()
		// Request Body
		reqBody := c.Request.Body

		// Log format
		logger.Infof("| %3d | %13v | %15s | %s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			reqBody,
		)
	}
}
