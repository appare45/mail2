package entity

import (
	"reflect"
	"strings"
	"testing"

	"github.com/appare45/mail2/smtp"
)

func TestParseEmail(t *testing.T) {
	tests := []struct {
		input string
		want  Email
	}{
		{
			input: "Smith@bar.com",
			want:  NewEmail("Smith", NewDomain("bar.com")),
		},
	}

	for _, test := range tests {
		email := NewEmail("", NewDomain(""))
		err := email.Parse(smtp.NewScanner(strings.NewReader(test.input)))
		t.Log(email.String())
		if err != nil {
			t.Errorf("ParseEmail(%q) returned error: %v", test.input, err)
		}
		if !reflect.DeepEqual(email, test.want) {
			t.Errorf("ParseEmail(%q) = %#v, want %#v", test.input, email, test.want)
		}
	}
}
