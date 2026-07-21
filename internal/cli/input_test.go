package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig_FromFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")

	data := []byte(`
server:
  host: localhost
`)

	err := os.WriteFile(path, data, 0644)
	require.NoError(t, err)

	opts := Options{
		Path: path,
	}

	cfg, err := LoadConfig(opts)

	require.NoError(t, err)

	server := cfg["server"].(map[string]any)

	assert.Equal(t, "localhost", server["host"])
}

func TestLoadConfig_FromStdin(t *testing.T) {
	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()

	file, err := os.CreateTemp("", "stdin")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.WriteString(`{"server":{"host":"localhost"}}`)
	require.NoError(t, err)

	_, err = file.Seek(0, 0)
	require.NoError(t, err)

	os.Stdin = file

	opts := Options{
		Stdin: true,
	}

	cfg, err := LoadConfig(opts)

	require.NoError(t, err)

	server := cfg["server"].(map[string]any)
	assert.Equal(t, "localhost", server["host"])
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	opts := Options{
		Path: "missing-config.yaml",
	}

	cfg, err := LoadConfig(opts)

	require.Error(t, err)
	assert.Nil(t, cfg)
}
