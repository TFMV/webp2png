package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	webp2png "github.com/TFMV/webp2png"
	"github.com/docopt/docopt-go"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal. Goodbye!")
		os.Exit(0)
	}()

	usage := `WebP to PNG Converter.

Usage:
  webp2png --input=<webp_file> --output=<png_file>
  webp2png -h | --help

Options:
  -h --help                Show this screen.
  --input=<webp_file>     Path to the input WebP file.
  --output=<png_file>     Path to the output PNG file.
`

	arguments, err := docopt.ParseDoc(usage)
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	webpPath, err := arguments.String("--input")
	if err != nil {
		log.Fatalf("Error getting input path: %v", err)
	}

	pngPath, err := arguments.String("--output")
	if err != nil {
		log.Fatalf("Error getting output path: %v", err)
	}

	// Validate input file exists
	if _, err := os.Stat(webpPath); os.IsNotExist(err) {
		log.Fatalf("Input file does not exist: %s", webpPath)
	}

	// Validate output directory exists
	outDir := filepath.Dir(pngPath)
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		log.Fatalf("Output directory does not exist: %s", outDir)
	}

	err = webp2png.ConvertWebPToPNG(webpPath, pngPath)
	if err != nil {
		log.Fatalf("Error converting image: %v", err)
	}

	fmt.Printf("Successfully converted %s to %s\n", webpPath, pngPath)
}
