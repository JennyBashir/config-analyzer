package analyzer

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"github.com/JennyBashir/config-analyzer/internal/rules"
	"github.com/JennyBashir/config-analyzer/internal/types"
)

func Analyze(cfg config.Config) []types.Issue {
	var issues []types.Issue
	issues = append(issues, rules.CheckDebug(cfg)...)
	issues = append(issues, rules.CheckPassword(cfg)...)
	issues = append(issues, rules.CheckTLS(cfg)...)
	issues = append(issues, rules.CheckAlgorithm(cfg)...)
	return issues
}
