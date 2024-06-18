package webp2png

import (
	"image/png"
	"os"

	"golang.org/x/image/webp"
)

// ConvertWebPToPNG reads a WebP file, decodes it, and writes it as a PNG file
func ConvertWebPToPNG(webpPath, pngPath string) error {
	// Open the WebP file
	webpFile, err := os.Open(webpPath)
	if err != nil {
		return err
	}
	defer webpFile.Close()

	// Decode the WebP image
	img, err := webp.Decode(webpFile)
	if err != nil {
		return err
	}

	// Create the PNG file
	pngFile, err := os.Create(pngPath)
	if err != nil {
		return err
	}
	defer pngFile.Close()

	// Encode the image to PNG and write it to the file
	err = png.Encode(pngFile, img)
	if err != nil {
		return err
	}

	return nil
}
