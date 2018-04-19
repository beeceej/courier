package main

import (
	"log"

	"github.com/beeceej/courier/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// address is the address to your Courier instance
	address = "localhost:8000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMailClient(conn)

	// Contact the server and print out its response.
	r, err := c.Send(context.Background(), &pb.SendBody{
		Recipients: []string{"email@email.com"},
		Sender:     "email@email.com",
		Subject:    "Hello, to my friends",
		BodyText:   "Just saying hi",
		BodyType:   "text/html",
	})
	if err != nil {
		log.Fatalf("could not send mail: %v", err)
	}
	log.Printf("\nCourier Server Response :: %s\n", r.GetStatus().String())
}
