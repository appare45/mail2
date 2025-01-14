package commands

import (
	"github.com/appare45/mail2/smtp"
	"github.com/appare45/mail2/smtp/entity"
)

type MailFrom struct {
	reversePath entity.Email
}

func (mailFrom MailFrom) Email() entity.Email {
	return mailFrom.reversePath
}

func NewMailFrom(reversePath entity.Email) MailFrom {
	return MailFrom{reversePath: reversePath}
}

func (mailFrom MailFrom) Command(conn *smtp.SmtpConnection) (*smtp.Response, error) {
	return conn.Cmd(221, "MAIL FROM: %s", mailFrom.reversePath)
}
