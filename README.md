# File Compressor

A modular, command-line file compression and decompression tool written in Go. This application currently supports the ZIP format, with a clean and extensible design that allows for easy addition of new compression formats, multi-threading, and other enhancements. It includes logging for debugging and configuration management for flexibility.

## Features

- **Compression**: Compress files into ZIP format.
- **Decompression**: Extract files from ZIP archives.
- **Modular Design**: Organized code structure for maintainability and extensibility.
- **Logging**: Built-in logging for tracking operations and errors.
- **Configuration**: Configurable settings (e.g., compression level) for future enhancements.
- **Extensibility**: Interface-based design to support additional formats (e.g., GZIP) and features like multi-threading.

## Project Structure

```bash
file-compressor/
├── cmd/
│   └── filecomp/
│       └── main.go           # Application entry point
├── internal/
│   ├── compressor/
│   │   ├── compressor.go    # Compression interface and factory
│   │   └── zip.go           # ZIP compression implementation
│   ├── decompressor/
│   │   ├── decompressor.go  # Decompression interface and factory
│   │   └── zip.go           # ZIP decompression implementation
│   ├── config/
│       └── config.go        # Configuration management
├── pkg/
│   └── logger/
│       └── logger.go        # Logging utility
├── go.mod                   # Go module definition
└── README.md                # Project documentation
```

## How It Works

### Command-Line Input
The user specifies the operation (`-compress` or `-decompress`), input file (`-input`), and output file (`-output`) via flags.

Example:
```bash
./filecomp -compress -input input.txt -output output.zip
```

### Initialization
- Initializes logger (pkg/logger).
- Loads configuration (internal/config).

### Compression
- Creates ZIP compressor implementing `Compressor` interface.
- `Compress()`:
    - Opens input file.
    - Creates ZIP archive.
    - Adds input file using Deflate algorithm.
    - Logs success or errors.

### Decompression
- Creates ZIP decompressor implementing `Decompressor` interface.
- `Decompress()`:
    - Opens ZIP archive.
    - Extracts files to output path.
    - Logs success or errors.

### Extensibility
- Interfaces allow adding formats (e.g., GZIP).
- Configuration system supports enhancements like compression level or multi-threading.

## Prerequisites

- **Go**: Version 1.18 or higher. Install from [golang.org](https://golang.org/).
- **Git**: Optional, for cloning the repository.

## Setup

### Clone the Repository
```bash
git clone https://github.com/yourusername/file-compressor.git
cd file-compressor
```

### Initialize Go Module
```bash
go mod init github.com/yourusername/file-compressor
go mod tidy
```

### Build the Application
```bash
go build -o filecomp ./cmd/filecomp
```

## Usage

### Compress a File
```bash
./filecomp -compress -input input.txt -output output.zip
```

### Decompress a File
```bash
./filecomp -decompress -input output.zip -output decompressed.txt
```

## Example Output
```bash
$ ./filecomp -compress -input input.txt -output output.zip
[FileCompressor] 2025/04/25 12:34:56 INFO: Compression completed successfully

$ ./filecomp -decompress -input output.zip -output decompressed.txt
[FileCompressor] 2025/04/25 12:35:00 INFO: Decompression completed successfully
```

## Error Handling

- Missing required flags show usage instructions.
- Invalid flag combinations result in an error.
- File I/O errors are logged and exit the program.

## Adding New Compression Formats

1. Create a new file in `internal/compressor` (e.g., `gzip.go`).
2. Implement the `Compressor` interface.
3. Extend the factory method.

Example:
```go
type gzipCompressor struct {
    log *logger.Logger
}

func (g *gzipCompressor) Compress(inputPath, outputPath string, cfg *config.Config) error {
    // Implement GZIP compression
}
```

## Future Enhancements

- Multi-threading
- Support for additional formats (GZIP, TAR, etc.)
- Custom compression levels
- Unit tests
- Progress reporting

## Contributing

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit changes: `git commit -m "Add my feature"`
4. Push to GitHub: `git push origin feature/my-feature`
5. Open a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For questions or feedback, open an issue or contact the maintainer at [jimitchavdadev@gmail.com](mailto:jimitchavdadev@gmail.com).

## Notes

- For Windows, use `filecomp.exe` instead of `./filecomp`.
- This tool works with any file type for ZIP compression/decompression.

