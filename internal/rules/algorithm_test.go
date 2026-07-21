package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/types"
	"testing"

	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckAlgorithm_WeakAlgorithms(t *testing.T) {
	keys := []string{
		"algorithm",
		"alg",
		"digest",
		"hash",
		"cipher",
	}
	weak := []string{
		"md5",
		"sha1",
		"des",
		"3des",
	}

	cfg := config.Config{}
	var issues []types.Issue
	for _, key := range keys {
		for _, algorithm := range weak {
			cfg = config.Config{
				key: algorithm,
			}
			issues = CheckAlgorithm(cfg)

			require.Len(t, issues, 1, "key: %s, algorithm: %s", key, algorithm)
			assert.Equal(t, "HIGH", issues[0].Severity, "key: %s, algorithm: %s", key, algorithm)
		}
	}
}

func TestCheckAlgorithm_SafeAlgorithms(t *testing.T) {
	sameKeys := []string{
		"algorithm",
		"alg",
		"digest",
		"hash",
		"cipher",
	}
	safeAlgorithms := []string{
		"sha256",
		"sha512",
		"aes",
		"aes256",
		"chacha20",
		"blake2",
	}
	cfg := config.Config{}
	var issues []types.Issue
	for _, key := range sameKeys {
		for _, algorithm := range safeAlgorithms {
			cfg = config.Config{
				key: algorithm,
			}
			issues = CheckAlgorithm(cfg)

			require.Empty(t, issues, "key: %s, algorithm: %s", key, algorithm)
		}
	}
}

func TestCheckAlgorithm_UnrelatedKey(t *testing.T) {
	sameAlgorithms := []string{
		"md5",
		"sha1",
		"des",
		"3des",
	}
	safeKeys := []string{
		"name",
		"version",
		"host",
		"port",
		"timeout",
		"username",
		"password",
		"tls",
	}
	cfg := config.Config{}
	var issues []types.Issue
	for _, key := range safeKeys {
		for _, algorithm := range sameAlgorithms {
			cfg = config.Config{
				key: algorithm,
			}
			issues = CheckAlgorithm(cfg)

			require.Empty(t, issues, "key: %s, algorithm: %s", key, algorithm)
		}
	}
}
