package mail

import (
	"github.com/go-mail/mail"
)

type Mail struct {
	Sender     string
	Recipients []string
	Subject    string
	BodyText   string
	BodyType   string

	SMTPHost     string
	SMTPPort     int
	SMTPPassword string
}

func (m *Mail) Send() {
	msg := mail.NewMessage()
	msg.SetHeader("From", m.Sender)
	msg.SetHeader("To", m.Recipients...)
	msg.SetHeader("Subject", m.Subject)
	msg.SetBody(m.BodyType, m.BodyText)

	d := mail.NewDialer(m.SMTPHost, m.SMTPPort, m.Sender, m.SMTPPassword)

	if err := d.DialAndSend(msg); err != nil {
		panic(err)
	}
}
