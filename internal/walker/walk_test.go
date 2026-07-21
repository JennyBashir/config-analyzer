package walker

import (
	"testing"

	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestWalk_SimpleMap(t *testing.T) {
	cfg := config.Config{
		"host": "localhost",
		"port": 8080,
	}

	values := make(map[string]any)

	Walk("", cfg, func(path string, value any) {
		values[path] = value
	})

	assert.Equal(t, "localhost", values["host"])
	assert.Equal(t, 8080, values["port"])
}

func TestWalk_NestedMap(t *testing.T) {
	cfg := config.Config{
		"server": map[string]any{
			"tls": map[string]any{
				"enabled": true,
			},
		},
	}

	values := make(map[string]any)

	Walk("", cfg, func(path string, value any) {
		values[path] = value
	})

	assert.Equal(t, true, values["server.tls.enabled"])
}

func TestWalk_Array(t *testing.T) {
	cfg := config.Config{
		"servers": []any{
			map[string]any{
				"host": "localhost",
			},
			map[string]any{
				"host": "example.com",
			},
		},
	}

	values := make(map[string]any)

	Walk("", cfg, func(path string, value any) {
		values[path] = value
	})

	assert.Equal(t, "localhost", values["servers[0].host"])
	assert.Equal(t, "example.com", values["servers[1].host"])
}

func TestWalk_EmptyConfig(t *testing.T) {
	cfg := config.Config{}

	visited := make(map[string]any)

	Walk("", cfg, func(path string, value any) {
		visited[path] = value
	})

	assert.Contains(t, visited, "")
	assert.Equal(t, cfg, visited[""])
}
