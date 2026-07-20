package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type Config map[string]any

func ReadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(path))

	return Parse(data, ext)
}

func Parse(data []byte, ext string) (Config, error) {
	var cfg Config

	switch ext {
	case ".json":
		err := json.Unmarshal(data, &cfg)
		return cfg, err

	case ".yaml", ".yml":
		err := yaml.Unmarshal(data, &cfg)
		return cfg, err

	default:
		return nil, fmt.Errorf("")
	}
}
