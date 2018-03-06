package courier

import (
	"context"
	"fmt"
	"strconv"

	"github.com/beeceej/courier/mail"
	"github.com/beeceej/courier/pb"
)

type Courier struct{}

func (s *Courier) Send(ctx context.Context, in *pb.SendBody) (*pb.Response, error) {
	fmt.Printf("Message Received :: %v\n", in)
	port, _ := strconv.ParseInt(SMTPServerPort, 10, 64)

	mailer := mail.Mail{
		Sender:       in.GetSender(),
		Recipients:   in.GetRecipients(),
		Subject:      in.GetSubject(),
		BodyText:     in.GetBodyText(),
		SMTPHost:     SMTPHost,
		SMTPPassword: SMTPPassword,
		SMTPPort:     int(port),
		BodyType:     in.GetBodyType(),
	}
	fmt.Println(mailer)
	mailer.Send()
	return &pb.Response{Status: pb.Status_Success}, nil
}
