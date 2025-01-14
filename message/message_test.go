package message

import (
	"testing"
	"time"

	"github.com/appare45/mail2/smtp/commands"
)

/*
https://datatracker.ietf.org/doc/html/rfc5322#appendix-A.1.1
A Message from One Person to Another with Simple Addressing

From: John Doe <jdoe@machine.example>
To: Mary Smith <mary@example.net> # TODO
Subject: Saying Hello # TODO
Date: Fri, 21 Nov 1997 09:55:06 -0600
Message-ID: <1234@local.machine.example> # TODO
*/
func TestMessageDataStream(t *testing.T) {
	from := mailbox_list("John Doe <jdoe@machine.example>")
	testTime := time.Date(1997, time.November, 21, 9, 55, 6, 0, time.FixedZone("CST", -6*60*60))
	date := NewDateTime(testTime)
	header := NewHeader(from, date)
	body := NewBody("This is a message just to say hello.\nSo, \"Hello\".")
	message := NewMessage(*header, body)

	expectedData := commands.Data_stream{
		"From: John Doe <jdoe@machine.example>",
		"Date: Fri, 21 Nov 1997 09:55:06 -0600",
		"",
		"This is a message just to say hello.",
		"So, \"Hello\".",
	}

	data := message.Data_stream()

	if len(data) != len(expectedData) {
		t.Fatalf("Expected data length to be %v, got %v", len(expectedData), len(data))
	}

	for i, line := range data {
		if line != expectedData[i] {
			t.Fatalf("Expected data line %v to be %v, got %v", i, expectedData[i], line)
		}
	}
}
