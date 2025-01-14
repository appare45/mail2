package entity

import (
	"fmt"
	"log/slog"
	"text/scanner"

	"github.com/appare45/mail2/smtp"
)

type Domain struct {
	string
}

func NewDomain(s string) Domain {
	return Domain{s}
}

func (d Domain) String() string {
	return string(d.string)
}

func (d *Domain) Parse(s smtp.Scanner) error {
	domain := ""
	defer func() {
		slog.Debug("Parsed", "domain", domain)
		d.string = domain
	}()
	for s.Peek() == ' ' { // Ignore leading spaces
		s.Next()
	}
	for s.Peek() != scanner.EOF {
		if s.Peek() == '\n' || s.Peek() == '\r' { // Stop at the end of the line
			if domain == "" {
				return fmt.Errorf("Domain is empty")
			}
			return nil
		}
		domain += string(s.Next())
	}
	if domain == "" {
		return fmt.Errorf("Domain is empty")
	}
	return nil
}
