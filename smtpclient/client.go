package smtpclient

type SmtpClient struct {
	SendEmail func(body EmailRequest) error
}
