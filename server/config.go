package server

import "github.com/david-allan-jones/contactserver/smtpclient"

type ServerConfig struct {
	UseHttps   bool
	TlsCert    string
	TlsKey     string
	SmtpClient smtpclient.SmtpClient
}
