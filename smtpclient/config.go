package smtpclient

type SmtpClientConfig struct {
	Sender   string
	Receiver string
	Server   string
	Port     int
	Password string
}
