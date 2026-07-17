package rules

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/JennyBashir/config-analyzer/internal/types"
)

func CheckDebug(cfg config.Config) []types.Issue {
	logValue, ok := cfg["log"]
	if !ok {
		return nil
	}

	logMap, ok := logValue.(map[string]any)
	if !ok {
		return nil
	}

	level, ok := logMap["level"].(string)
	if !ok {
		return nil
	}

	if level == "debug" {
		return nil
	}

	return []types.Issue{
		{
			Severity:       "LOW",
			Message:        "",
			Recommendation: "",
		},
	}
}
