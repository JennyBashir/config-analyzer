package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/JennyBashir/config-analyzer/internal/types"
	"github.com/JennyBashir/config-analyzer/internal/walker"
	"strings"
)

var suspiciousWords = []string{
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

func CheckPassword(cfg config.Config) []types.Issue {
	var issues []types.Issue

	walker.Walk("", cfg, func(path string, value any) {

		parts := strings.Split(path, ".")
		key := parts[len(parts)-1]

		str, ok := value.(string)
		if !ok {
			return
		}

		if containsKeyword(key, suspiciousWords) {
			if strings.HasPrefix(str, "$") || strings.HasPrefix(str, "${") {
				return
			}
			issues = append(issues, types.Issue{
				Severity:       "HIGH",
				Message:        "секрет хранится в открытом виде в: " + path,
				Recommendation: "используйте секрет-хранилище или переменные окружения",
			})
		}
	})
	return issues
}
