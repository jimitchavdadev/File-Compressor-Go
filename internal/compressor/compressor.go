package compressor

import (
	"github.com/jimitchavdadev/file-compressor/internal/config"
	"github.com/jimitchavdadev/file-compressor/pkg/logger"
)

type Compressor interface {
	Compress(inputPath, outputPath string, cfg *config.Config) error
}

func NewZIPCompressor(log *logger.Logger) Compressor {
	return &zipCompressor{log: log}
}
