package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/JennyBashir/config-analyzer/internal/types"
	"github.com/JennyBashir/config-analyzer/internal/walker"
	"strings"
)

func CheckTLS(cfg config.Config) []types.Issue {
	var issues []types.Issue

	walker.Walk("", cfg, func(path string, value any) {
		val, ok := value.(bool)
		if !ok {
			return
		}

		bad := false
		var message string
		var recommendation string

		if strings.HasSuffix(path, ".tls.enabled") && !val {
			message = "TLS отключен: " + path
			recommendation = "Включите TLS."
			bad = true
		}
		if (strings.HasSuffix(path, ".insecure_skip_verify") ||
			strings.HasSuffix(path, ".skip_tls_verify")) && val {
			message = "Отключена проверка TLS-сертификата: " + path
			recommendation = "Включите проверку TLS-сертификата."
			bad = true
		}

		if bad {
			issues = append(issues, types.Issue{
				Severity:       "HIGH",
				Message:        message,
				Recommendation: recommendation,
			})
		}
	})
	return issues
}
