package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jimitchavdadev/file-compressor/internal/compressor"
	"github.com/jimitchavdadev/file-compressor/internal/config"
	"github.com/jimitchavdadev/file-compressor/internal/decompressor"
	"github.com/jimitchavdadev/file-compressor/pkg/logger"
)

func main() {
	// Define command-line flags
	compress := flag.Bool("compress", false, "Compress the input file")
	decompress := flag.Bool("decompress", false, "Decompress the input file")
	input := flag.String("input", "", "Input file path")
	output := flag.String("output", "", "Output file path")
	flag.Parse()

	// Validate flags
	if *input == "" || *output == "" {
		fmt.Println("Error: input and output file paths are required")
		flag.Usage()
		os.Exit(1)
	}
	if !(*compress || *decompress) {
		fmt.Println("Error: specify either -compress or -decompress")
		flag.Usage()
		os.Exit(1)
	}

	// Initialize logger
	log := logger.NewLogger()

	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error("Failed to load configuration: %v", err)
		os.Exit(1)
	}

	// Perform compression or decompression
	if *compress {
		comp := compressor.NewZIPCompressor(log)
		if err := comp.Compress(*input, *output, cfg); err != nil {
			log.Error("Compression failed: %v", err)
			os.Exit(1)
		}
		log.Info("Compression completed successfully")
	} else {
		decomp := decompressor.NewZIPDecompressor(log)
		if err := decomp.Decompress(*input, *output, cfg); err != nil {
			log.Error("Decompression failed: %v", err)
			os.Exit(1)
		}
		log.Info("Decompression completed successfully")
	}
}
