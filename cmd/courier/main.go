package main

import (
	"log"
	"net"

	"github.com/beeceej/courier"
	"github.com/beeceej/courier/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var (
		lis net.Listener
		err error
	)
	if lis, err = net.Listen("tcp", courier.CourierPort); err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	} else {
		log.Printf("Listening on %v\n", lis.Addr())
	}

	s := grpc.NewServer()
	pb.RegisterMailServer(s, &courier.Courier{})

	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
