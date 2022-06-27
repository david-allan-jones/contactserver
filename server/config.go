package server

import "github.com/david-allan-jones/contactserver/smtpclient"

type ServerConfig struct {
	Port       int
	Path       string
	TlsCert    string
	TlsKey     string
	SmtpClient smtpclient.SmtpClient
}
