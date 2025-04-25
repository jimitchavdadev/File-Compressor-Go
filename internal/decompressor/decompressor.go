package decompressor

import (
	"github.com/jimitchavdadev/file-compressor/internal/config"
	"github.com/jimitchavdadev/file-compressor/pkg/logger"
)

type Decompressor interface {
	Decompress(inputPath, outputPath string, cfg *config.Config) error
}

func NewZIPDecompressor(log *logger.Logger) Decompressor {
	return &zipDecompressor{log: log}
}
