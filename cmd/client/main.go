package main

import (
	"context"
	pb "github.com/JennyBashir/config-analyzer/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewConfigAnalyzerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config := `
server: 
host: 0.0.0.0
`
	responce, err := client.Analyze(ctx, &pb.AnalyzeRequest{
		Config: config,
	})
	if err != nil {
		log.Fatalf("Analyze failed: %v", err)
	}

	for _, issue := range responce.Issues {
		log.Printf(
			"\n%s:\n%s\nRecommendation: %s\n",
			issue.Severity,
			issue.Message,
			issue.Recommendation,
		)
	}
}
