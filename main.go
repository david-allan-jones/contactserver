package main

import (
	"fmt"

	"github.com/david-allan-jones/contactserver/appConfig"
	"github.com/david-allan-jones/contactserver/server"
	"github.com/david-allan-jones/contactserver/smtpclient"
)

func main() {
	appConfig, err := appConfig.Create()
	if err != nil {
		fmt.Printf("Error creating app configuration: %v\n", err)
		return
	}
	client := smtpclient.CreateSmtpClient(&smtpclient.SmtpClientConfig{
		Sender:   appConfig.SmtpSender,
		Receiver: appConfig.SmtpReceiver,
		Server:   appConfig.SmtpServer,
		Port:     appConfig.SmtpPort,
		Password: appConfig.Password,
	})
	server.Start(server.ServerConfig{
		Port:       appConfig.HttpServerPort,
		Path:       appConfig.HttpPath,
		TlsCert:    appConfig.TlsCert,
		TlsKey:     appConfig.TlsKey,
		SmtpClient: client,
	})
}
