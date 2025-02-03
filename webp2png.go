package web2png

import (
	"errors"
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/webp"
)

// Common errors returned by the package
var (
	ErrEmptyPath     = errors.New("empty file path provided")
	ErrInvalidInput  = errors.New("invalid input file")
	ErrInvalidOutput = errors.New("invalid output path")
	ErrDecoding      = errors.New("failed to decode WebP")
	ErrEncoding      = errors.New("failed to encode PNG")
)

// ConvertWebPToPNG converts a WebP image file to PNG format.
// It takes the input WebP file path and desired output PNG file path as parameters.
// Returns an error if the conversion fails due to invalid paths, read/write errors,
// or image conversion issues.
func ConvertWebPToPNG(webpPath, pngPath string) error {
	// Validate input parameters
	if webpPath == "" || pngPath == "" {
		return ErrEmptyPath
	}

	// Ensure input file exists and is readable
	if _, err := os.Stat(webpPath); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	// Ensure output directory exists
	outDir := filepath.Dir(pngPath)
	if _, err := os.Stat(outDir); err != nil {
		return fmt.Errorf("%w: output directory does not exist", ErrInvalidOutput)
	}

	// Open input file
	webpFile, err := os.Open(webpPath)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}
	defer webpFile.Close()

	// Decode WebP image
	img, err := webp.Decode(webpFile)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDecoding, err)
	}

	// Create output file
	pngFile, err := os.Create(pngPath)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidOutput, err)
	}
	defer pngFile.Close()

	// Encode to PNG
	if err := png.Encode(pngFile, img); err != nil {
		return fmt.Errorf("%w: %v", ErrEncoding, err)
	}

	return nil
}
