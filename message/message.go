package message

import (
	"strings"

	"github.com/appare45/mail2/smtp/commands"
)

// Internet Message Format
// https://datatracker.ietf.org/doc/html/rfc5322
type Message struct {
	header Header
	body   Body
}

type Body string

func NewBody(body string) Body {
	return Body(body)
}

func (b Body) split() []string {
	return strings.Split(string(b), "\n")
}

func (b Body) Data_stream() commands.Data_stream {
	return b.split()
}

func NewMessage(header Header, body Body) *Message {
	return &Message{header: header, body: body}
}

func (m *Message) Data_stream() commands.Data_stream {
	var data_stream commands.Data_stream
	// Construct Header
	data_stream = append(data_stream, m.header.Data_stream()...)
	// Separater between Header and Body
	data_stream = append(data_stream, "")
	data_stream = append(data_stream, m.body.Data_stream()...)
	return data_stream
}
