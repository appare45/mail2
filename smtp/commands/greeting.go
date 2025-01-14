package commands

import (
	"github.com/appare45/mail2/smtp"
	"github.com/appare45/mail2/smtp/entity"
)

type Greeting struct {
	domain entity.Domain
}

func NewGreeting(domain entity.Domain) *Greeting {
	return &Greeting{domain: domain}
}

func (g *Greeting) Response() *smtp.Response {
	return smtp.NewResponse(220, g.domain.String())
}
