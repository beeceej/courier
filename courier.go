package courier

import (
	"context"
	"fmt"

	"github.com/beeceej/courier/pb"
)

type Courier struct{}

func (s *Courier) Send(ctx context.Context, in *pb.SendBody) (*pb.Response, error) {
	fmt.Printf("Message Received :: %v\n", in)
	return &pb.Response{Status: pb.Status_Success}, nil
}
