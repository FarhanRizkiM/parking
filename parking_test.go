// parking_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEnterParking(t *testing.T) {
	// Setup
	router := gin.Default()
	router.POST("/enter", enterParking)

	// Perform request
	req, _ := http.NewRequest("POST", "/enter", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Entered the parking")
}

func TestExitParking(t *testing.T) {
	// Setup
	router := gin.Default()
	router.POST("/exit", exitParking)

	// Perform request
	req, _ := http.NewRequest("POST", "/exit", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Exited the parking")
}

func TestGenerateQRCode(t *testing.T) {
	// Setup
	router := gin.Default()
	router.GET("/generate-qr", generateQR)

	// Perform request
	req, _ := http.NewRequest("GET", "/generate-qr", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "QR code generated successfully")
}
