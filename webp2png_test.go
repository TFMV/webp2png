package web2png

import (
	"embed"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/foo.webp
var testdata embed.FS

func TestConvertWebPToPNG(t *testing.T) {
	// Create temp test directory
	tmpDir, err := os.MkdirTemp("", "webp2png-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Setup test fixtures
	testWebPData, err := testdata.ReadFile("testdata/foo.webp")
	require.NoError(t, err)

	testWebP := filepath.Join(tmpDir, "test.webp")
	require.NoError(t, os.WriteFile(testWebP, testWebPData, 0644))

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
			webpPath: testWebP,
			pngPath:  filepath.Join(tmpDir, "out2.png"),
			setupFunc: func() error {
				return os.WriteFile(testWebP, []byte("invalid webp data"), 0644)
			},
			expectedErr: ErrDecoding,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test case
			if tt.setupFunc != nil {
				require.NoError(t, tt.setupFunc())
			} else if tt.expectedErr != ErrEmptyPath && tt.expectedErr != ErrInvalidInput {
				// Restore valid WebP data for non-error cases
				require.NoError(t, os.WriteFile(testWebP, testWebPData, 0644))
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

func TestConvertWebPToPNG_RealFiles(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// Skip if test file doesn't exist
	webpPath := filepath.Join("testdata", "sample.webp")
	if _, err := os.Stat(webpPath); os.IsNotExist(err) {
		t.Skip("skipping test: sample.webp not found in testdata/")
	}

	tmpDir, err := os.MkdirTemp("", "webp2png-integration-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	pngPath := filepath.Join(tmpDir, "output.png")
	err = ConvertWebPToPNG(webpPath, pngPath)
	assert.NoError(t, err)

	stat, err := os.Stat(pngPath)
	require.NoError(t, err)
	assert.Greater(t, stat.Size(), int64(0), "PNG file should not be empty")
}
