package courier

import (
	"context"

	"github.com/beeceej/courier/mail"
	"github.com/beeceej/courier/pb"
)

// Courier is the server object for the courier grpc interface
type Courier struct{}

// Send is the rpc call to send mail
func (s *Courier) Send(ctx context.Context, in *pb.SendBody) (*pb.Response, error) {
	mailer := mail.Mail{
		Message: *mail.New(
			mail.CC(in.GetCc()...),
			mail.BCC(in.GetBcc()...),
			mail.To(in.GetRecipients()...),
			mail.From(in.GetSender()),
			mail.Subject(in.GetSubject()),
			mail.Body(in.GetBodyText()),
			mail.BodyType(in.GetBodyType()),
			mail.Attachments(in.GetAttachments()...),
		),
		SMTP: mail.SMTP{
			Host:     SMTPHost,
			Password: SMTPPassword,
			Port:     int(SMTPServerPort),
		},
	}
	mailer.Send()
	return &pb.Response{Status: pb.Status_Success}, nil
}
