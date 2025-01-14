package message

import "github.com/appare45/mail2/smtp/commands"

type FieldName string
type FieldBody string

func NewFieldBody(body string) (FieldBody, error) {
	return FieldBody(body), nil
}

/*
*

	Header fields are lines beginning with a field name, followed by a
	colon (":"), followed by a field body, and terminated by CRLF.  A
	field name MUST be composed of printable US-ASCII characters (i.e.,
	characters that have values between 33 and 126, inclusive), except
	colon.  A field body may be composed of printable US-ASCII characters
	as well as the space (SP, ASCII value 32) and horizontal tab (HTAB,
	ASCII value 9) characters (together known as the white space
	characters, WSP).  A field body MUST NOT include CR and LF except
	when used in "folding" and "unfolding", as described in section
	2.2.3.  All field bodies MUST conform to the syntax described in
	sections 3 and 4 of this specification.
*/
type Header struct {
	from      mailbox_list
	orig_date date_time
}

type mailbox_list string

func (l mailbox_list) String() string {
	return string(l)
}

func NewHeader(from mailbox_list, orig_date date_time) *Header {
	return &Header{
		from:      from,
		orig_date: orig_date,
	}
}

const MAX_LENGTH = 998

func (h *Header) Data_stream() commands.Data_stream {
	var data commands.Data_stream
	data = append(data, ("From: " + h.from.String()))
	data = append(data, ("Date: " + h.orig_date.String()))
	return data
}
