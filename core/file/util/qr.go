package util

import (
	"image/png"
	"os"
	"path/filepath"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// QRCodeType represents the type of data to encode in the QR code.
type QRCodeType int

const (
	TextQRCode QRCodeType = iota
	URLQRCode
	LocationQRCode
	PhoneNumberQRCode
	EmailQRCode
	// Add more QR code types here as needed
)

const (
	QRCodeCorrectionLevelLow      = qr.L // 7% of codewords can be restored.
	QRCodeCorrectionLevelMedium   = qr.M // 15% of codewords can be restored.
	QRCodeCorrectionLevelQuartile = qr.Q // 25% of codewords can be restored.
	QRCodeCorrectionLevelHigh     = qr.H // 30% of codewords can be restored.
)

// GenerateQRCode generates a QR code from the given data and saves it to the "public/storage" directory.
//
// Arguments:
// - data: the data to encode in the QR code.
// - filename: the name of the output file to save the QR code.
// - errorCorrectionLevel: the error correction level to use for the QR code.
// - encoding: the encoding to use for the QR code.
//
// Returns:
// - string: the path to the output file.
// - error: an error if the QR code generation or file saving failed.
func GenerateQRCode(data string, filename string, errorCorrectionLevel qr.ErrorCorrectionLevel, encoding qr.Encoding) (string, error) {
	// Generate the QR code image
	qrCode, err := qr.Encode(data, errorCorrectionLevel, encoding)
	if err != nil {
		return "", err
	}

	// Scale the QR code image to 256x256 pixels
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		return "", err
	}

	// Open the output file for writing
	outputPath := filepath.Join("public", "storage", filename)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	// Encode the QR code matrix as a PNG image and write it to the output file
	err = png.Encode(outputFile, qrCode)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
