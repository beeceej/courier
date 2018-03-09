package mail

// SMTP represents the details of the SMTP server we're speaking with
type SMTP struct {
	Host     string
	Port     int
	Password string
}
