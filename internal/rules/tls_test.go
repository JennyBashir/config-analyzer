package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckTLS_Disabled(t *testing.T) {
	cfg := config.Config{
		"tls": map[string]any{
			"enabled": false,
		},
	}
	issues := CheckTLS(cfg)
	require.Len(t, issues, 1)
	assert.Equal(t, "HIGH", issues[0].Severity)

	cfg = config.Config{
		"server": map[string]any{
			"tls": map[string]any{
				"enabled": false,
			},
		},
	}
	issues = CheckTLS(cfg)
	require.Len(t, issues, 1)
	assert.Equal(t, "HIGH", issues[0].Severity)
}

func TestCheckTLS_Enabled(t *testing.T) {
	cfg := config.Config{
		"tls": map[string]any{
			"enabled": true,
		},
	}
	issues := CheckTLS(cfg)
	require.Empty(t, issues)

	cfg = config.Config{
		"server": map[string]any{
			"tls": map[string]any{
				"enabled": true,
			},
		},
	}
	issues = CheckTLS(cfg)
	require.Empty(t, issues)
}

func TestCheckTLS_SkipVerify(t *testing.T) {
	cfg := config.Config{
		"tls": map[string]any{
			"insecure_skip_verify": true,
		},
	}
	issues := CheckTLS(cfg)
	require.Len(t, issues, 1)
	assert.Equal(t, "HIGH", issues[0].Severity)

	cfg = config.Config{
		"server": map[string]any{
			"tls": map[string]any{
				"skip_tls_verify": true,
			},
		},
	}
	issues = CheckTLS(cfg)
	require.Len(t, issues, 1)
	assert.Equal(t, "HIGH", issues[0].Severity)
}

func TestCheckTLS_VerificationEnabled(t *testing.T) {
	cfg := config.Config{
		"tls": map[string]any{
			"insecure_skip_verify": false,
		},
	}
	issues := CheckTLS(cfg)
	require.Empty(t, issues)

	cfg = config.Config{
		"server": map[string]any{
			"tls": map[string]any{
				"skip_tls_verify": false,
			},
		},
	}
	issues = CheckTLS(cfg)
	require.Empty(t, issues)
}
