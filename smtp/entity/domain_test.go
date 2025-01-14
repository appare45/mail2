package entity

import (
	"reflect"
	"strings"
	"testing"

	"github.com/appare45/mail2/smtp"
)

func TestParseDomain(t *testing.T) {
	var tests = []struct {
		input    string
		expected Domain
	}{
		{"test.com", NewDomain("test.com")},
		{"test.com\r\n", NewDomain("test.com")},
		{"test.com\n", NewDomain("test.com")},
	}
	for _, test := range tests {
		domain := NewDomain("")
		scanner := smtp.NewScanner(strings.NewReader(test.input))
		got := domain.Parse(scanner)
		if got != nil {
			t.Errorf("ParseDomain(%q) = %v", test.input, got)
		}
		if !reflect.DeepEqual(domain, test.expected) {
			t.Errorf("ParseDomain(%q) = %#v (Expected %#v)", test.input, domain, test.expected)
		}
	}
}
