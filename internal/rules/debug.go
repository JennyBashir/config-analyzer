package rules

import (
	"strings"

	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/JennyBashir/config-analyzer/internal/types"
	"github.com/JennyBashir/config-analyzer/internal/walker"
)

func CheckDebug(cfg config.Config) []types.Issue {
	var issues []types.Issue

	walker.Walk("", cfg, func(path string, value any) {
		level, ok := value.(string)
		if !ok {
			return
		}

		if path == "log.level" && strings.EqualFold(level, "debug") {
			issues = append(issues, types.Issue{
				Severity:       "LOW",
				Message:        "включено логирование в debug-режиме: " + path,
				Recommendation: "Используйте уровень логирования info или выше в production.",
			})
		}
	})

	return issues
}
