package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestParse_JSON(t *testing.T) {
	data := []byte(`{
"debug": true,
"port": 8080
}`)

	cfg, err := Parse(data, ".json")

	require.NoError(t, err)
	assert.Equal(t, true, cfg["debug"])
	assert.Equal(t, float64(8080), cfg["port"])
}

func TestParse_YAML(t *testing.T) {
	data := []byte(`
debug: true
port: 8080
`)

	cfg, err := Parse(data, ".yaml")

	require.NoError(t, err)
	assert.Equal(t, true, cfg["debug"])
	assert.Equal(t, 8080, cfg["port"])
}

func TestParse_InvalidJSON(t *testing.T) {
	data := []byte(`{
		"debug": true,
		"port":
	}`)

	cfg, err := Parse(data, ".json")

	require.Error(t, err)
	assert.Nil(t, cfg)
}

func TestParse_InvalidYAML(t *testing.T) {
	data := []byte(`
debug: true
port: [
`)

	cfg, err := Parse(data, ".yaml")

	require.Error(t, err)
	assert.Nil(t, cfg)
}

func TestParse_UnsupportedExtension(t *testing.T) {
	data := []byte("test")

	cfg, err := Parse(data, ".txt")

	require.Error(t, err)
	assert.Nil(t, cfg)
}

func TestReadConfig_JSON(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "config.json")

	data := []byte(`{
		"debug": true
	}`)

	require.NoError(t, os.WriteFile(path, data, 0644))

	cfg, err := ReadConfig(path)

	require.NoError(t, err)
	assert.Equal(t, true, cfg["debug"])
}

func TestReadConfig_YAML(t *testing.T) {
	dir := t.TempDir()

	path := filepath.Join(dir, "config.yaml")

	data := []byte(`
debug: true
port: 8080
`)

	require.NoError(t, os.WriteFile(path, data, 0644))

	cfg, err := ReadConfig(path)

	require.NoError(t, err)
	assert.Equal(t, true, cfg["debug"])
	assert.Equal(t, 8080, cfg["port"])
}

func TestReadConfig_FileNotFound(t *testing.T) {
	cfg, err := ReadConfig("does-not-exist.yaml")

	require.Error(t, err)
	assert.Nil(t, cfg)
}
