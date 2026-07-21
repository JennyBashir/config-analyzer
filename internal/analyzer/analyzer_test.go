package analyzer

import (
	"testing"

	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestAnalyze_SafeConfig(t *testing.T) {
	cfg := config.Config{
		"log": map[string]any{
			"level": "info",
		},
		"server": map[string]any{
			"host": "127.0.0.1",
			"tls": map[string]any{
				"enabled": true,
			},
		},
		"database": map[string]any{
			"password": "${DB_PASSWORD}",
		},
		"security": map[string]any{
			"algorithm": "sha256",
		},
	}

	issues := Analyze(cfg)

	assert.Empty(t, issues)
}

func TestAnalyze_DebugIssue(t *testing.T) {
	cfg := config.Config{
		"log": map[string]any{
			"level": "debug",
		},
	}

	issues := Analyze(cfg)

	assert.Len(t, issues, 1)
	assert.Equal(t, "LOW", issues[0].Severity)
}

func TestAnalyze_PasswordIssue(t *testing.T) {
	cfg := config.Config{
		"database": map[string]any{
			"password": "123456",
		},
	}

	issues := Analyze(cfg)

	assert.Len(t, issues, 1)
	assert.Equal(t, "HIGH", issues[0].Severity)
}

func TestAnalyze_TLSIssue(t *testing.T) {
	cfg := config.Config{
		"server": map[string]any{
			"tls": map[string]any{
				"enabled": false,
			},
		},
	}

	issues := Analyze(cfg)

	assert.Len(t, issues, 1)
	assert.Equal(t, "HIGH", issues[0].Severity)
}

func TestAnalyze_OpenBindIssue(t *testing.T) {
	cfg := config.Config{
		"server": map[string]any{
			"host": "0.0.0.0",
		},
	}

	issues := Analyze(cfg)

	assert.Len(t, issues, 1)
	assert.Equal(t, "MEDIUM", issues[0].Severity)
}

func TestAnalyze_AlgorithmIssue(t *testing.T) {
	cfg := config.Config{
		"security": map[string]any{
			"algorithm": "md5",
		},
	}

	issues := Analyze(cfg)

	assert.Len(t, issues, 1)
	assert.Equal(t, "HIGH", issues[0].Severity)
}

func TestAnalyze_MultipleIssues(t *testing.T) {
	cfg := config.Config{
		"log": map[string]any{
			"level": "debug",
		},
		"server": map[string]any{
			"host": "0.0.0.0",
			"tls": map[string]any{
				"enabled": false,
			},
		},
		"database": map[string]any{
			"password": "123456",
		},
		"security": map[string]any{
			"algorithm": "md5",
		},
	}

	issues := Analyze(cfg)

	assert.Len(t, issues, 5)
}
