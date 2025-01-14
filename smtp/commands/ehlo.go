package commands

import (
	"github.com/appare45/mail2/smtp"
	"github.com/appare45/mail2/smtp/entity"
)

type Ehlo struct {
	domain entity.Domain
}

func (e Ehlo) Domain() entity.Domain {
	return e.domain
}

func NewEhlo(domain entity.Domain) Ehlo {
	return Ehlo{domain: domain}
}

func (e Ehlo) Command(conn *smtp.SmtpConnection) (*smtp.Response, error) {
	return conn.Cmd(250, "EHLO %s", e.domain)
}
