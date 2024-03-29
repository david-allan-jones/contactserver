package appConfig

import (
	"errors"
	"os"
	"strconv"
)

type AppConfig struct {
	SmtpSender   string
	SmtpReceiver string
	SmtpServer   string
	SmtpPort     int
	Password     string
	UseHttps     bool
	TlsCert      string
	TlsKey       string
}

func Create() (*AppConfig, error) {
	smtpPort, smtpPortErr := strconv.Atoi(os.Getenv("CONTACTSERVER_SMTP_PORT"))
	if smtpPortErr != nil {
		return nil, errors.New("CONTACTSERVER_SMTP_PORT env variable needs to be a number")
	}

	useHttps, useHttpsErr := strconv.ParseBool(os.Getenv("CONTACTSERVER_USE_HTTPS"))
	if useHttpsErr != nil {
		return nil, errors.New("USE_HTTPS env variable needs to be a boolean")
	}

	return &AppConfig{
		SmtpSender:   os.Getenv("CONTACTSERVER_SMTP_SENDER"),
		SmtpReceiver: os.Getenv("CONTACTSERVER_SMTP_RECEIVER"),
		SmtpServer:   os.Getenv("CONTACTSERVER_SMTP_SERVER"),
		SmtpPort:     smtpPort,
		Password:     os.Getenv("CONTACTSERVER_SMTP_PASSWORD"),
		UseHttps:     useHttps,
		TlsCert:      os.Getenv("CONTACTSERVER_TLS_CERT"),
		TlsKey:       os.Getenv("CONTACTSERVER_TLS_KEY"),
	}, nil
}
