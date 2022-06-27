package appConfig

import (
	"errors"
	"os"
	"strconv"
)

type AppConfig struct {
	SmtpSender     string
	SmtpReceiver   string
	SmtpServer     string
	SmtpPort       int
	Password       string
	HttpServerPort int
	HttpPath       string
}

func Create() (*AppConfig, error) {
	smtpPort, smtpPortErr := strconv.Atoi(os.Getenv("CONTACTSERVER_SMTP_PORT"))
	if smtpPortErr != nil {
		return nil, errors.New("CONTACTSERVER_SMTP_PORT env variable needs to be a number")
	}

	httpPort, httpPortErr := strconv.Atoi(os.Getenv("CONTACTSERVER_HTTP_PORT"))
	if httpPortErr != nil {
		return nil, errors.New("CONTACTSERVER_HTTP_PORT env variable needs to be a number")
	}

	return &AppConfig{
		SmtpSender:     os.Getenv("CONTACTSERVER_SMTP_SENDER"),
		SmtpReceiver:   os.Getenv("CONTACTSERVER_SMTP_RECEIVER"),
		SmtpServer:     os.Getenv("CONTACTSERVER_SMTP_SERVER"),
		SmtpPort:       smtpPort,
		Password:       os.Getenv("CONTACTSERVER_SMTP_PASSWORD"),
		HttpPath:       os.Getenv("CONTACTSERVER_HTTP_PATH"),
		HttpServerPort: httpPort,
	}, nil
}
