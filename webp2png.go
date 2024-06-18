package webp2png

import (
	"image/png"
	"os"

	"golang.org/x/image/webp"
)

func ConvertWebPToPNG(webpPath, pngPath string) error {
	webpFile, err := os.Open(webpPath)
	if err != nil {
		return err
	}
	defer webpFile.Close()

	img, err := webp.Decode(webpFile)
	if err != nil {
		return err
	}

	pngFile, err := os.Create(pngPath)
	if err != nil {
		return err
	}
	defer pngFile.Close()

	err = png.Encode(pngFile, img)
	if err != nil {
		return err
	}

	return nil
}
