package config

type Config struct {
	CompressionLevel int
}

func NewConfig() (*Config, error) {
	return &Config{
		CompressionLevel: 6, // Default ZIP compression level
	}, nil
}
