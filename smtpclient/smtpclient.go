package smtpclient

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func CreateSmtpClient(config *SmtpClientConfig) SmtpClient {
	client := SmtpClient{
		SendEmail: func(req EmailRequest) error {
			m := gomail.NewMessage()

			m.SetHeader("From", config.Sender)

			m.SetHeader("To", config.Receiver)
			m.SetHeader("Subject", "New submission from Contact Server")
			m.SetBody("text/html", fmt.Sprintf(`
				<!DOCTYPE html>
				<html>
				<head>
					<style>
						span {
							font-weight: bold;
						}
					</style>
				</head>
				<body>
					<p><span>Name: </span>%v</p>
					<p><span>Contact: </span>%v</p>
					<p><span>Message: </span></p>
					<p>%v</p>
				</body>
				</html>
			`, req.Name, req.Contact, req.Message))

			d := gomail.NewDialer(config.Server, int(config.Port), config.Sender, config.Password)

			d.TLSConfig = &tls.Config{
				InsecureSkipVerify: false,
				ServerName:         config.Server,
			}

			err := d.DialAndSend(m)
			if err != nil {
				fmt.Println("Failed sending email.")
				fmt.Println(err)
				return err
			}
			fmt.Println("Successfully sent email!")
			return nil
		},
	}
	return client
}
