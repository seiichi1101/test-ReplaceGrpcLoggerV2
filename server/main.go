package main

import (
	"context"
	"log"
	"net"

	pb "test-ReplaceGrpcLoggerV2/proto"

	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello World"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	zapLogger, err := zap.NewDevelopment()
	grpcZap.ReplaceGrpcLoggerV2(zapLogger)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcZap.UnaryServerInterceptor(zapLogger),
		),
	)
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
