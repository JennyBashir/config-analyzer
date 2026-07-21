package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JennyBashir/config-analyzer/internal/analyzer"
	"github.com/JennyBashir/config-analyzer/internal/cli"
)

func main() {
	opts, err := cli.ParseFlags()
	if err != nil {
		log.Fatalf("parse flags error: %v", err)
	}

	cfg, err := cli.LoadConfig(opts)
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

	if len(issues) > 0 && !opts.Silent {
		os.Exit(1)
	}
}
