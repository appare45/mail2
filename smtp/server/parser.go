package server

import (
	"fmt"
	"io"
	"log/slog"
	"text/scanner"

	"github.com/appare45/mail2/smtp"
	"github.com/appare45/mail2/smtp/commands"
	"github.com/appare45/mail2/smtp/entity"
)

type Handler interface {
	Data(commands.Data) *smtp.Response
	Ehlo(commands.Ehlo) *smtp.Response
	Rest(commands.Rset) *smtp.Response
	Quit(commands.Quit) *smtp.Response
	MailFrom(commands.MailFrom) *smtp.Response
	RcptTo(commands.RcptTo) *smtp.Response
}

/*
行単位でコマンドをパースする
TODO: Parse以上のことをしたいので名前を変える
*/
func Parse(r io.Reader, h Handler) (*smtp.Response, error) {
	s := smtp.NewScanner(r)
	defer func() {
		if s.Peek() == '\r' {
			s.Next()
		} else {
			return
		}
		if s.Peek() == '\n' {
			// println(scanner.TokenString(s.Next()))
			// s.Next()
		} else {
			return
		}
	}()
	var rawCommand string
	// 最初の4文字を見て何のコマンドが送られているのかを判断する
	for i := 0; i < 4; i++ {
		r := s.Next()
		if r == scanner.EOF {
			return nil, fmt.Errorf("EOF")
		}
		rawCommand += string(r)
	}
	slog.Debug("rawCommand: %s", "COMMAND", rawCommand)
	if s.Peek() == '\r' {
		switch rawCommand {
		case "DATA":
			return h.Data(commands.NewData(nil)), nil
		case "RSET":
			return h.Rest(*commands.NewRset()), nil
		case "QUIT":
			return h.Quit(commands.NewQuit()), nil
		}
	}
	switch rawCommand {
	case "EHLO", "HELO":
		domain := entity.NewDomain("")
		err := domain.Parse(s)
		if err != nil {
			return nil, err
		}
		return h.Ehlo(commands.NewEhlo(domain)), nil
	case "MAIL":
		for s.Peek() == ' ' {
			s.Next()
		}
		const FROM string = "FROM"
		for i := 0; i < len(FROM); i++ {
			r := s.Next()
			if r == scanner.EOF || r != rune(FROM[i]) {
				return nil, fmt.Errorf("EOF")
			}
		}
		for s.Peek() == ' ' {
			s.Next()
		}
		email := entity.NewEmail("", entity.NewDomain(""))
		if err := email.Parse(s); err != nil {
			return nil, err
		}
		return h.MailFrom(commands.NewMailFrom(email)), nil
	case "RCPT":
		for s.Peek() == ' ' {
			s.Next()
		}
		const TO = "TO"
		for i := 0; i < len(TO); i++ {
			r := s.Next()
			if r == scanner.EOF || r != rune(TO[i]) {
				return nil, fmt.Errorf("EOF")
			}
		}
		for s.Peek() == ' ' {
			s.Next()
		}
		email := entity.NewEmail("", entity.NewDomain(""))
		if err := email.Parse(s); err != nil {
			return nil, err
		}
		return h.RcptTo(commands.NewRcptTo(email)), nil
	}
	// TODO
	return h.Data(commands.NewData(nil)), nil
}
