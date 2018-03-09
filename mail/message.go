package mail

import (
	gomail "github.com/go-mail/mail"
)

// Msg is represents data which can be sent in an email
type Msg struct {
	// CC is a list of addresses who will be carbon copied
	CC []string

	// BCC is a list of addresses who will be blind carbon copied
	BCC []string

	// Recipients is a list of addresses who will be sent the mail
	Recipients []string

	// From is who the mail is from
	From string

	// Subject is what's listed as the subject on the message
	Subject string

	//Body is the message body
	Body string

	// BodyType is the MIME type of the message
	BodyType string

	// Attachments represents files to attach (looked up by filepath)
	Attachments []string
}

// MsgOption is a configuration for sending
type MsgOption func(m *Msg)

func New(options ...MsgOption) *Msg {
	m := new(Msg)
	for _, o := range options {
		o(m)
	}
	return m
}

// CC sets the addresses to CC
func CC(addresses ...string) MsgOption {
	return func(m *Msg) {
		m.CC = addresses
	}
}

// BCC sets the addresses to BCC
func BCC(addresses ...string) MsgOption {
	return func(m *Msg) {
		m.BCC = addresses
	}
}

// To sets the addresses to BCC
func To(addresses ...string) MsgOption {
	return func(m *Msg) {
		m.Recipients = addresses
	}
}

// From sets the message property to the address the message is from
func From(address string) MsgOption {
	return func(m *Msg) {
		m.From = address
	}
}

// Body sets the message body
func Body(body string) MsgOption {
	return func(m *Msg) {
		m.Body = body
	}
}

// BodyType sets the message content type
func BodyType(bodyType string) MsgOption {
	return func(m *Msg) {
		m.BodyType = bodyType
	}
}

// Subject sets the message content type
func Subject(subject string) MsgOption {
	return func(m *Msg) {
		m.Subject = subject
	}
}

// Attachments sets the message content type
func Attachments(paths ...string) MsgOption {
	return func(m *Msg) {
		m.Attachments = paths
	}
}

// CreateGoMail creates a go-mail message
func (m *Msg) CreateGoMail() *gomail.Message {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", m.Recipients...)
	for _, cc := range m.CC {
		msg.SetAddressHeader("Cc", cc, cc)
	}
	msg.SetHeader("Subject", m.Subject)
	msg.SetBody(m.BodyType, m.Body)
	return msg
}
