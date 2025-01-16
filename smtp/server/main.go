package server

import (
	"fmt"
	"net"

	"github.com/appare45/mail2/smtp"
	"github.com/appare45/mail2/smtp/commands"
	"github.com/appare45/mail2/smtp/entity"
)

type Server struct {
	port   string
	domain entity.Domain
}

// TODO: Validate port
func NewServer(port string, domain entity.Domain) *Server {
	return &Server{port, domain}
}

func (s Server) Start() {
	l, err := net.Listen("tcp", s.port)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go func(c net.Conn) {
			defer c.Close()
			connection := smtp.IntoSmtpConnection(c)
			connection.WriteResponse(commands.NewGreeting(s.domain).Response())
			h := NewSmtpHandler()
			for {
				// ここでレスポンスまで作って帰ってきていてほしい
				cmd, err := Parse(c, h)
				if err != nil {
					fmt.Printf("failed to parse command: %v", err)
					break
				}
				connection.WriteResponse(cmd)
			}
		}(conn)
	}
}
