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
		cfg, err = config.Parse(data, ".json")
		if err != nil {
			cfg, err = config.Parse(data, ".yaml")
			if err != nil {
				return nil, err
			}
		}
	} else {
		var err error
		cfg, err = config.ReadConfig(opts.Path)
		if err != nil {
			return nil, err
		}
	}
	return cfg, nil
}
