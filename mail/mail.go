package mail

import (
	"github.com/go-mail/mail"
)

type Mail struct {
	Message Msg
	SMTP    SMTP
}

func (m *Mail) Send() {
	msg := m.Message.CreateGoMail()

	d := mail.NewDialer(m.SMTP.Host, m.SMTP.Port, m.Message.From, m.SMTP.Password)

	if err := d.DialAndSend(msg); err != nil {
		panic(err)
	}
}
