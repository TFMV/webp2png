package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TFMV/webp2png/internal/converter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	outDir  string
	verbose bool
	rootCmd = &cobra.Command{
		Use:   "webp2png [file.webp]",
		Short: "Convert WebP images to PNG format",
		Long: `A robust WebP to PNG converter.
Automatically saves the PNG file to the same directory as the WebP file
unless an output directory is specified.`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			webpPath := args[0]

			// Validate input file exists
			if _, err := os.Stat(webpPath); os.IsNotExist(err) {
				return fmt.Errorf("input file does not exist: %s", webpPath)
			}

			// Determine output path
			baseName := filepath.Base(webpPath)
			pngName := replaceExt(baseName, ".png")

			var pngPath string
			if outDir != "" {
				// Use specified output directory
				if _, err := os.Stat(outDir); os.IsNotExist(err) {
					return fmt.Errorf("output directory does not exist: %s", outDir)
				}
				pngPath = filepath.Join(outDir, pngName)
			} else {
				// Use same directory as input file
				pngPath = filepath.Join(filepath.Dir(webpPath), pngName)
			}

			// Convert the image
			if err := converter.ConvertWebPToPNG(webpPath, pngPath); err != nil {
				return fmt.Errorf("conversion failed: %w", err)
			}

			if verbose {
				fmt.Printf("Successfully converted %s to %s\n", webpPath, pngPath)
			}

			return nil
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.webp2png.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outDir, "outdir", "o", "", "output directory (default is same as input)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	viper.BindPFlag("outdir", rootCmd.PersistentFlags().Lookup("outdir"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".webp2png" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".webp2png")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// replaceExt replaces the file extension
func replaceExt(filename, newExt string) string {
	ext := filepath.Ext(filename)
	return filename[:len(filename)-len(ext)] + newExt
}

func main() {
	Execute()
}
