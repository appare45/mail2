package server

import (
	"github.com/appare45/mail2/smtp"
	"github.com/appare45/mail2/smtp/commands"
	"github.com/appare45/mail2/smtp/entity"
)

// WIP
type SmtpHandler struct {
	serverHostName entity.Domain
	rcptTo         entity.Email
	mailFrom       entity.Email
}

func NewSmtpHandler() SmtpHandler {
	return SmtpHandler{}
}

func (handler SmtpHandler) Data(c commands.Data) *smtp.Response {
	return smtp.NewResponse(200, "OK")
}

func (handler SmtpHandler) MailFrom(c commands.MailFrom) *smtp.Response {
	handler.mailFrom = c.Email()
	return smtp.NewResponse(200, "OK")
}

func (handler SmtpHandler) RcptTo(c commands.RcptTo) *smtp.Response {
	handler.rcptTo = c.Email()
	return smtp.NewResponse(200, "OK")
}

func (handler SmtpHandler) Rest(c commands.Rset) *smtp.Response {
	handler = SmtpHandler{}
	return smtp.NewResponse(200, "OK")
}

func (handler SmtpHandler) Quit(c commands.Quit) *smtp.Response {
	return smtp.NewResponse(200, "OK")
}

func (handler SmtpHandler) Ehlo(c commands.Ehlo) *smtp.Response {
	handler.serverHostName = c.Domain()
	return smtp.NewResponse(200, "OK")
}
