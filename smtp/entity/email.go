package entity

import (
	"fmt"
	"text/scanner"

	"github.com/appare45/mail2/smtp"
)

type Email struct {
	local  string
	domain Domain
}

func (e Email) String() string {
	return e.local + "@" + e.domain.String()
}

func (e Email) Local() string {
	return e.local
}

func (e Email) Domain() Domain {
	return e.domain
}

func NewEmail(local string, domain Domain) Email {
	return Email{local: local, domain: domain}
}

func (e *Email) Parse(s smtp.Scanner) error {
	for s.Peek() != '@' && s.Peek() != scanner.EOF {
		e.local += string(s.Next())
	}

	if s.Next() != '@' {
		return fmt.Errorf("expected '@' but got %q", s.Peek())
	}

	if err := e.domain.Parse(s); err != nil {
		return err
	}

	return nil
}
