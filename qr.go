// qr.go
package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/skip2/go-qrcode"
)

func generateQRCode(data string) {
	q, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	file, err := os.Create("qrcode.png")
	if err != nil {
		fmt.Println("Error creating QR code file:", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, q.Image(256))
	if err != nil {
		fmt.Println("Error encoding QR code to PNG:", err)
		return
	}
	fmt.Println("QR code generated successfully")
}
