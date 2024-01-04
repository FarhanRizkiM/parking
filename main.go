// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes for parking application
	router.POST("/enter", enterParking)
	router.POST("/exit", exitParking)
	router.GET("/generate-qr", generateQR)

	router.Run(":8080")
}
