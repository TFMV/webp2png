# WebP to PNG Converter

A command-line tool for converting WebP images to PNG format.

## Features

- Simple command-line interface
- Automatic output path generation
- Custom output directory support
- Configuration via config file, environment variables, or command-line flags
- Graceful error handling

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/TFMV/webp2png.git
cd webp2png

# Build the binary
go build -o webp2png ./cmd

# Install to your PATH (optional)
sudo mv webp2png /usr/local/bin/
```

### For Apple Silicon (M1/M2/M3) Macs

```bash
GOOS=darwin GOARCH=arm64 go build -o webp2png ./cmd
sudo mv webp2png /usr/local/bin/
```

## Usage

### Basic Usage

Convert a WebP file to PNG in the same directory:

```bash
webp2png image.webp
```

This will create `image.png` in the same directory as the input file.

### Specify Output Directory

```bash
webp2png --outdir=/path/to/output image.webp
```

This will create the PNG file in the specified output directory.

### Verbose Output

```bash
webp2png --verbose image.webp
```

### Using a Config File

Create a `.webp2png.yaml` file in your home directory:

```yaml
outdir: "/path/to/default/output"
verbose: true
```

Or specify a custom config file:

```bash
webp2png --config=myconfig.yaml image.webp
```

## Command-Line Options

```bash
Usage:
  webp2png [file.webp]
  webp2png -h | --help

Options:
  -h --help                Show this screen.
  --outdir, -o             Output directory (default is same as input).
  --verbose, -v            Verbose output.
  --config                 Config file (default is $HOME/.webp2png.yaml).
```

## Environment Variables

The tool also supports configuration via environment variables:

```bash
export WEBP2PNG_OUTDIR="/path/to/output"
export WEBP2PNG_VERBOSE="true"
```

## Development

### Prerequisites

- Go 1.18 or higher
- Required dependencies:
  - github.com/spf13/cobra
  - github.com/spf13/viper
  - golang.org/x/image/webp

### Running Tests

```bash
go test -v ./...
```

### Project Structure

```bash
webp2png/
├── cmd/
│   └── main.go            # Command-line interface
├── internal/
│   └── converter/         # Core conversion logic
│       ├── converter.go
│       └── converter_test.go
├── testdata/              # Test files
│   └── foo.webp
├── go.mod
├── go.sum
└── README.md
```

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
