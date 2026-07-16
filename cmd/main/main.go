package main

import (
	"github.com/JennyBashir/config-analyzer/internal/config"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: checker <config-file>")
	}

	path := os.Args[1]

	conf, err := config.ParseConfig(path)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Print("success config loading ", conf)
}
