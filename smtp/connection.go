package smtp

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/textproto"
)

type SmtpConnection struct {
	rowConn net.Conn
	conn    *textproto.Conn
	reader  io.Reader
	scanner *bufio.Scanner
}

/**
 * WriteResponse writes a response to the connection
 * @param response: Response to write
 * @return error: Error if any
 */
// TODO: Responseに対してWriteを実装したほうがまともな気がする
func (c SmtpConnection) WriteResponse(response *Response) error {
	separater := " "
	if response.followNext {
		separater = "-"
	}
	return c.Write("%d%s%s", response.code, separater, response.text)
}

/**
 * Write writes a raw message to the connection
 * @param format: Format of the message
 * @param args: Arguments to the format
 * @return error: Error if any
 */
func (c SmtpConnection) Write(format string, args ...any) error {
	slog.Debug("Sending", "MESSAGE", fmt.Sprintf(format, args...))
	_, error := c.conn.Cmd(format, args...)
	return error
}

func (c SmtpConnection) Read(expectCode int) (*Response, error) {
	code, message, error := c.conn.ReadResponse(expectCode)
	slog.Debug("Received", "CODE", code, "MESSAGE", message)
	return &Response{code: code, text: message}, error
}

func (c SmtpConnection) Cmd(expectCode int, format string, args ...any) (*Response, error) {
	error := c.Write(format, args...)
	if error != nil {
		return &Response{}, error
	}
	return c.Read(expectCode)
}

func (c *SmtpConnection) Scanner() *bufio.Scanner {
	return c.scanner
}

func NewSmtpConnection(ipaddr net.TCPAddr) (*SmtpConnection, error) {
	conn, err := net.Dial("tcp", ipaddr.String())
	slog.Debug("Connected", "ip", ipaddr.String())
	if err != nil {
		return nil, err
	}
	return IntoSmtpConnection(conn), nil
}

func IntoSmtpConnection(conn net.Conn) *SmtpConnection {
	textProtoConn := textproto.NewConn(conn)
	bufreader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(bufreader)
	return &SmtpConnection{
		conn,
		textProtoConn,
		bufreader,
		scanner,
	}
}

func (c *SmtpConnection) Close() error {
	if c.conn == nil {
		return fmt.Errorf("Connection is not established")
	}
	err := (*c.conn).Close()
	if err != nil {
		return err
	}
	return nil
}
