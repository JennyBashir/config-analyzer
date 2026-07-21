package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckOpenBind_OpenAddress(t *testing.T) {
	keys := []string{
		"host",
		"listen",
		"bind",
		"address",
		"addr",
	}

	for _, key := range keys {
		cfg := config.Config{
			key: "0.0.0.0",
		}
		issues := CheckOpenBind(cfg)

		require.Len(t, issues, 1)
		assert.Equal(t, "MEDIUM", issues[0].Severity)
	}
}

func TestCheckOpenBind_LocalAddresses(t *testing.T) {
	addresses := []string{
		"127.0.0.1",
		"localhost",
		"192.168.1.10",
		"10.0.0.1",
		"::1",
	}
	sameKeys := []string{
		"host",
		"listen",
		"bind",
		"address",
		"addr",
	}

	for _, key := range sameKeys {
		for _, address := range addresses {
			cfg := config.Config{
				key: address,
			}
			issues := CheckOpenBind(cfg)

			require.Empty(t, issues)
		}
	}
}

func TestCheckOpenBind_UnrelatedKey(t *testing.T) {
	otherKeys := []string{
		"name",
		"version",
		"timeout",
		"username",
		"password",
		"port",
		"tls",
		"debug",
	}

	for _, key := range otherKeys {
		cfg := config.Config{
			key: "0.0.0.0",
		}
		issues := CheckOpenBind(cfg)

		require.Empty(t, issues)
	}
}
