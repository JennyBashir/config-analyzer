package rules

import (
	"strings"

	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/JennyBashir/config-analyzer/internal/types"
	"github.com/JennyBashir/config-analyzer/internal/walker"
)

var weakAlgorithms = []string{
	"md5",
	"sha1",
	"des",
	"3des",
}
var algorithmKeys = []string{
	"algorithm",
	"alg",
	"digest",
	"hash",
	"cipher",
}

func CheckAlgorithm(cfg config.Config) []types.Issue {
	var issues []types.Issue

	walker.Walk("", cfg, func(path string, value any) {

		parts := strings.Split(path, ".")
		key := parts[len(parts)-1]

		str, ok := value.(string)
		if !ok {
			return
		}
		str = strings.ToLower(str)

		if containsKeyword(key, algorithmKeys) {
			for _, word := range weakAlgorithms {
				if str == word {
					issues = append(issues, types.Issue{
						Severity:       "HIGH",
						Message:        "используется небезопасный алгоритм: " + str + " (" + path + ")",
						Recommendation: "Используйте современные алгоритмы, например SHA-256.",
					})
				}
			}
		}
	})
	return issues
}
