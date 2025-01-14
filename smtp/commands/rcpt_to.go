package commands

import (
	"github.com/appare45/mail2/smtp"
	"github.com/appare45/mail2/smtp/entity"
)

type RcptTo struct {
	forwardPath entity.Email
}

func (rcptTo RcptTo) Email() entity.Email {
	return rcptTo.forwardPath
}

func NewRcptTo(forwardPath entity.Email) RcptTo {
	return RcptTo{forwardPath: forwardPath}
}

func (rcptTo RcptTo) Command(conn *smtp.SmtpConnection) (*smtp.Response, error) {
	return conn.Cmd(221, "RCPT TO: <%s>", rcptTo.forwardPath)
}
