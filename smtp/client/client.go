package client

import (
	"net"

	"github.com/appare45/mail2/smtp"
	"github.com/appare45/mail2/smtp/commands"
)

type SmtpClient struct {
	connection *smtp.SmtpConnection
}

func NewSmtpClient(ip net.TCPAddr) (*SmtpClient, error) {
	conn, err := smtp.NewSmtpConnection(ip)
	if err != nil {
		return nil, err
	}
	_, error := conn.Read(220)
	if error != nil {
		return nil, error
	}
	c := &SmtpClient{
		connection: conn,
	}
	return c, nil
}

func (client SmtpClient) Command(command commands.SmtpCommand) (*smtp.Response, error) {
	return command.Command(client.connection)
}

func (c *SmtpClient) Close() {
	c.connection.Close()
}
