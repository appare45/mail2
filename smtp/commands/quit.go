package commands

import (
	"github.com/appare45/mail2/smtp"
)

type Quit struct{}

func NewQuit() Quit {
	return Quit{}
}

func (q Quit) Command(conn *smtp.SmtpConnection) (*smtp.Response, error) {
	return conn.Cmd(221, "QUIT")
}
