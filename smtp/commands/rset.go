package commands

import (
	"github.com/appare45/mail2/smtp"
)

type Rset struct{}

func NewRset() *Rset {
	return &Rset{}
}

func (r Rset) Command(conn *smtp.SmtpConnection) (*smtp.Response, error) {
	return conn.Cmd(250, "RSET")
}
