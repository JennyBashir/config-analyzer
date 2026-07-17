package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/JennyBashir/config-analyzer/internal/types"
	"strconv"
	"strings"
)

var suspiciousWords = []string{
	"pass",
	"pwd",
	"secret",
	"token",
	"apikey",
	"api_key",
	"private",
}

func isSuspiciousKey(key string) bool {
	key = strings.ToLower(key)

	for _, word := range suspiciousWords {
		if strings.Contains(key, word) {
			return true
		}
	}
	return false
}

func findPasswords(path string, value any) []types.Issue {
	var issues []types.Issue

	switch v := value.(type) {

	case map[string]any:
		var nextPath string

		for key, child := range v {
			if path == "" {
				nextPath = key
			} else {
				nextPath = path + "." + key
			}
			if isSuspiciousKey(key) {
				if _, ok := child.(string); ok {
					issues = append(issues, types.Issue{
						Severity:       "HIGH",
						Message:        "пароль хранится в открытом виде: " + nextPath,
						Recommendation: "используйте секрет-хранилище или переменные окружения",
					})
				}
			}
			issues = append(issues, findPasswords(nextPath, child)...)
		}

	case []any:
		var nextPath string
		for ind, val := range v {
			if path == "" {
				nextPath = "[" + strconv.Itoa(ind) + "]"
			} else {
				nextPath = path + "[" + strconv.Itoa(ind) + "]"
			}
			issues = append(issues, findPasswords(nextPath, val)...)
		}
	}
	return issues
}

func CheckPassword(cfg config.Config) []types.Issue {
	return findPasswords("", cfg)
}
