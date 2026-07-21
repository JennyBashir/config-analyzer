package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckPassword_PlaintextSecret(t *testing.T) {
	suspicious := []string{
		"pass",
		"pwd",
		"secret",
		"token",
		"apikey",
		"api_key",
		"password",
		"passwd",
		"access_token",
		"auth_token",
		"client_secret",
	}

	for _, key := range suspicious {
		cfg := config.Config{
			key: "my-secret-password",
		}
		issues := CheckPassword(cfg)

		require.Len(t, issues, 1)
		assert.Equal(t, "HIGH", issues[0].Severity)
	}
}

func TestCheckPassword_EnvironmentVariables(t *testing.T) {
	values := []string{
		"$PASSWORD",
		"${PASSWORD}",
		"$API_KEY",
		"${CLIENT_SECRET}",
	}
	suspicious := []string{
		"pass",
		"pwd",
		"secret",
		"token",
		"apikey",
		"api_key",
		"password",
		"passwd",
		"access_token",
		"auth_token",
		"client_secret",
	}

	for _, key := range suspicious {
		for _, value := range values {
			cfg := config.Config{
				key: value,
			}
			issues := CheckPassword(cfg)

			require.Empty(t, issues)
		}
	}
}

func TestCheckPassword_UnrelatedKey(t *testing.T) {
	keys := []string{
		"name",
		"version",
		"host",
		"port",
		"timeout",
		"algorithm",
		"log",
		"tls",
	}

	for _, key := range keys {
		cfg := config.Config{
			key: "my-secret-password",
		}
		issues := CheckPassword(cfg)

		require.Empty(t, issues)
	}
}
