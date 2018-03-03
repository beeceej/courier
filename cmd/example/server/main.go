package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/beeceej/courier/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// should be env variable
const port = ":50000"

type server struct{}

func (s *server) Send(ctx context.Context, in *pb.SendMailRequest) (*pb.SendStatus, error) {
	fmt.Printf("Message Received :: %v\n", in)
	return &pb.SendStatus{Status: pb.SendStatus_Success}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMailServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
