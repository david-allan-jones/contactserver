package server

import "github.com/david-allan-jones/contactserver/smtpclient"

type ServerConfig struct {
	Port       int
	Path       string
	UseHttps   bool
	TlsCert    string
	TlsKey     string
	SmtpClient smtpclient.SmtpClient
}
