package main

import (
	"fmt"
	"github.com/JennyBashir/config-analyzer/internal/analyzer"
	"github.com/JennyBashir/config-analyzer/internal/config"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: checker <config-file>")
	}

	path := os.Args[1]

	cfg, err := config.ParseConfig(path)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	issues := analyzer.Analyze(cfg)

	for _, issue := range issues {
		fmt.Printf(
			"%s: %s\nRecommendation: %s\n\n",
			issue.Severity,
			issue.Message,
			issue.Recommendation,
		)
	}

	if len(issues) > 0 {
		os.Exit(1)
	}
}
