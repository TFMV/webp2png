package webp2png

import (
	"os"
	"testing"
)

func TestConvertWebPToPNG(t *testing.T) {
	webpPath := "foo.webp"
	pngPath := "foo.png"

	err := ConvertWebPToPNG(webpPath, pngPath)
	if err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	if _, err := os.Stat(pngPath); os.IsNotExist(err) {
		t.Fatalf("PNG file was not created")
	} else {
		t.Logf("PNG file created at: %s", pngPath)
	}
}
