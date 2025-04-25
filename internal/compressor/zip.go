package compressor

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/jimitchavdadev/file-compressor/internal/config"
	"github.com/jimitchavdadev/file-compressor/pkg/logger"
)

type zipCompressor struct {
	log *logger.Logger
}

func (z *zipCompressor) Compress(inputPath, outputPath string, cfg *config.Config) error {
	// Create output ZIP file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Create ZIP writer
	zw := zip.NewWriter(outFile)
	defer zw.Close()

	// Open input file
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file info
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Create ZIP header
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filepath.Base(inputPath)
	header.Method = zip.Deflate

	// Write file to ZIP
	writer, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
