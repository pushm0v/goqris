package main

import (
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"os"
)

func main() {
	// open and decode image file
	file, _ := os.Open("qris.jpg")
	img, _, _ := image.Decode(file)

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	fmt.Println(result)

	qrWriter := qrcode.NewQRCodeWriter()
	newQr, _ := qrWriter.Encode(result.String(), gozxing.BarcodeFormat_QR_CODE, 400, 400, nil)
	file, _ = os.Create("qris_new.jpg")
	defer file.Close()

	// *BitMatrix implements the image.Image interface,
	// so it is able to be passed to jpeg.Encode directly.
	_ = jpeg.Encode(file, newQr, nil)
}
