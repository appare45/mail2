package smtp

import (
	"io"
	"text/scanner"
)

type Scanner interface {
	Peek() rune
	Next() rune
}

type smtpScanner struct {
	r *scanner.Scanner
}

func NewScanner(r io.Reader) Scanner {
	scanner := scanner.Scanner{}
	scanner.Init(r)
	return &smtpScanner{r: &scanner}
}

// TODO: Add logger
func (s *smtpScanner) Peek() rune {
	return s.r.Peek()
}

func (s *smtpScanner) Next() rune {
	r := s.r.Next()
	// slog.Debug("Read", "rune", scanner.TokenString(r))
	return r
}
