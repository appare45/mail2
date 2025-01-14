package commands

import "github.com/appare45/mail2/smtp"

type SmtpCommand interface {
	Command(conn *smtp.SmtpConnection) (*smtp.Response, error)
}
