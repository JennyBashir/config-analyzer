package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	pb "github.com/JennyBashir/config-analyzer/gen"
	grpcserver "github.com/JennyBashir/config-analyzer/internal/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterConfigAnalyzerServer(
		server,
		&grpcserver.Server{},
	)

	go func() {
		log.Println("gRPC server started on :50051")

		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("shutting down gRPC server")

	server.GracefulStop()
}
