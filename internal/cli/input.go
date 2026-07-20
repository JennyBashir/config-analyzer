package cli

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"io"
	"os"
)

func LoadConfig(opts Options) (config.Config, error) {
	var cfg config.Config

	if opts.Stdin {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}

		extSlice := []string{
			".json",
			".yaml",
			".yml",
		}

		for _, ext := range extSlice {
			cfg, err = config.Parse(data, ext)
			if err == nil {
				return cfg, nil
			}
		}

		return nil, err
	}
	var err error
	cfg, err = config.ReadConfig(opts.Path)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
