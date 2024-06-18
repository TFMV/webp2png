package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/TFMV/webp2png"
)

func main() {
	webpPath := flag.String("input", "", "Path to the input WebP file")
	pngPath := flag.String("output", "", "Path to the output PNG file")
	flag.Parse()

	if *webpPath == "" || *pngPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := webp2png.ConvertWebPToPNG(*webpPath, *pngPath)
	if err != nil {
		log.Fatalf("Error converting image: %v", err)
	}

	fmt.Printf("Successfully converted %s to %s\n", *webpPath, *pngPath)
}
