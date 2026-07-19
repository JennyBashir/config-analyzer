package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/JennyBashir/config-analyzer/internal/types"
	"github.com/JennyBashir/config-analyzer/internal/walker"
	"strings"
)

var suspiciousList = []string{
	"host",
	"listen",
	"bind",
	"address",
	"addr",
}

func containsKeyword(key string, list []string) bool {
	key = strings.ToLower(key)

	for _, word := range list {
		if strings.Contains(key, word) {
			return true
		}
	}
	return false
}

func CheckOpenBind(cfg config.Config) []types.Issue {
	var issues []types.Issue

	walker.Walk("", cfg, func(path string, value any) {

		parts := strings.Split(path, ".")
		key := parts[len(parts)-1]

		str, ok := value.(string)
		if !ok {
			return
		}

		if containsKeyword(key, suspiciousList) {
			if str == "0.0.0.0" {
				issues = append(issues, types.Issue{
					Severity:       "MEDIUM",
					Message:        "Использование `0.0.0.0` без ограничений в: " + path,
					Recommendation: "Используйте конкретный IP вместо 0.0.0.0.",
				})
			}
		}
	})
	return issues
}
