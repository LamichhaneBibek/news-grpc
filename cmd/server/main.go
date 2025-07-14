package main

import (
	"log"
	"net"

	newsv1 "github.com/LamichhaneBibek/news-grpc/api/news/v1"
	ingrpc "github.com/LamichhaneBibek/news-grpc/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	newsv1.RegisterNewsServiceServer(s, ingrpc.NewServer())
	healthSrv := health.NewServer()
	healthv1.RegisterHealthServer(s, healthSrv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
