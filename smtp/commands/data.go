package commands

import (
	"github.com/appare45/mail2/smtp"
)

type Data_stream []string

func (d Data_stream) Get() []string {
	return d
}

type Data struct {
	message Data_stream
}

func NewData(message Data_stream) Data {
	return Data{
		message,
	}
}

func (d Data) Command(conn *smtp.SmtpConnection) (*smtp.Response, error) {
	err := conn.Write("DATA")
	if err != nil {
		return nil, err
	}
	for _, line := range d.message.Get() {
		err = conn.Write(line)
		if err != nil {
			return nil, err
		}
	}
	conn.Write(".")
	if err != nil {
		return nil, err
	}
	return conn.Read(250)
}
