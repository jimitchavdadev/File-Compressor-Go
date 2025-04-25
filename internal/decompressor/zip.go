package decompressor

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/jimitchavdadev/file-compressor/internal/config"
	"github.com/jimitchavdadev/file-compressor/pkg/logger"
)

type zipDecompressor struct {
	log *logger.Logger
}

func (z *zipDecompressor) Decompress(inputPath, outputPath string, cfg *config.Config) error {
	// Open ZIP file
	reader, err := zip.OpenReader(inputPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Check if ZIP contains a single file
	if len(reader.File) == 1 {
		// For single-file ZIPs, use the provided outputPath directly
		file := reader.File[0]
		zf, err := file.Open()
		if err != nil {
			return err
		}
		defer zf.Close()

		// Create output file at the specified outputPath
		outFile, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		// Copy file contents
		_, err = io.Copy(outFile, zf)
		return err
	}

	// For multiple files, extract to a directory
	for _, file := range reader.File {
		zf, err := file.Open()
		if err != nil {
			return err
		}
		defer zf.Close()

		// Use outputPath as a directory and preserve original file names
		outPath := filepath.Join(outputPath, file.Name)
		// Ensure the directory exists
		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}

		// Create output file
		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer outFile.Close()

		// Copy file contents
		_, err = io.Copy(outFile, zf)
		if err != nil {
			return err
		}
	}
	return nil
}
