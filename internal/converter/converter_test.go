package converter

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertWebPToPNG(t *testing.T) {
	// Create temp test directory
	tmpDir, err := os.MkdirTemp("", "webp2png-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Use the existing test file
	testWebP := "/Users/thomasmcgeehan/web2png/web2png/testdata/foo.webp"

	// Verify the test file exists
	_, err = os.Stat(testWebP)
	if os.IsNotExist(err) {
		t.Skip("Test file not found: /Users/thomasmcgeehan/web2png/web2png/testdata/foo.webp")
	}
	require.NoError(t, err)

	tests := []struct {
		name        string
		webpPath    string
		pngPath     string
		setupFunc   func() error
		expectedErr error
	}{
		{
			name:        "empty paths",
			webpPath:    "",
			pngPath:     "",
			expectedErr: ErrEmptyPath,
		},
		{
			name:        "non-existent input",
			webpPath:    "nonexistent.webp",
			pngPath:     filepath.Join(tmpDir, "out.png"),
			expectedErr: ErrInvalidInput,
		},
		{
			name:        "non-existent output directory",
			webpPath:    testWebP,
			pngPath:     filepath.Join("nonexistent", "out.png"),
			expectedErr: ErrInvalidOutput,
		},
		{
			name:     "successful conversion",
			webpPath: testWebP,
			pngPath:  filepath.Join(tmpDir, "out.png"),
		},
		{
			name:     "invalid webp data",
			webpPath: filepath.Join(tmpDir, "invalid.webp"),
			pngPath:  filepath.Join(tmpDir, "out2.png"),
			setupFunc: func() error {
				return os.WriteFile(filepath.Join(tmpDir, "invalid.webp"), []byte("invalid webp data"), 0644)
			},
			expectedErr: ErrDecoding,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test case
			if tt.setupFunc != nil {
				require.NoError(t, tt.setupFunc())
			}

			// Run test
			err := ConvertWebPToPNG(tt.webpPath, tt.pngPath)

			// Verify results
			if tt.expectedErr != nil {
				assert.True(t, errors.Is(err, tt.expectedErr),
					"expected error %v, got %v", tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				// Verify output file exists and is not empty
				stat, err := os.Stat(tt.pngPath)
				assert.NoError(t, err, "output file should exist")
				assert.Greater(t, stat.Size(), int64(0), "output file should not be empty")
			}
		})
	}
}
