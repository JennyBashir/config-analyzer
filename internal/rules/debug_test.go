package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckDebug_DebugEnabled(t *testing.T) {
	cfg := config.Config{
		"log": map[string]any{
			"level": "debug",
		},
	}

	issues := CheckDebug(cfg)

	require.Len(t, issues, 1)
	assert.Equal(t, "LOW", issues[0].Severity)
}

func TestCheckDebug_SafeLevels(t *testing.T) {
	levels := []string{
		"info",
		"warn",
		"warning",
		"error",
		"fatal",
	}

	for _, level := range levels {
		cfg := config.Config{
			"log": map[string]any{
				"level": level,
			},
		}
		issues := CheckDebug(cfg)

		require.Empty(t, issues)
	}
}

func TestCheckDebug_UnrelatedPath(t *testing.T) {
	configs := []config.Config{
		{
			"app": map[string]any{
				"level": "debug",
			},
		},
		{
			"database": map[string]any{
				"level": "debug",
			},
		},
		{
			"server": map[string]any{
				"level": "debug",
			},
		},
		{
			"logger": map[string]any{
				"mode": "debug",
			},
		},
		{
			"logging": map[string]any{
				"level": "debug",
			},
		},
		{
			"log": map[string]any{
				"mode": "debug",
			},
		},
		{
			"level": "debug",
		},
		{
			"debug": "debug",
		},
	}

	for _, cfg := range configs {
		issues := CheckDebug(cfg)

		require.Empty(t, issues)
	}
}
