package grpc

import newsv1 "github.com/LamichhaneBibek/news-grpc/api/news/v1"

type Server struct {
	newsv1.UnimplementedNewsServiceServer
}

func NewServer() *Server {
	return &Server{}
}
