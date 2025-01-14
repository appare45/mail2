package smtp

import (
	"testing"
)

func TestResponse(t *testing.T) {
	tests := []struct {
		code int
		text string
	}{
		{220, "foo.com Simple Mail Transfer Service Ready"},
		{550, "No such user here"},
		{354, "Start mail input; end with <CRLF>.<CRLF>"},
	}

	for _, tt := range tests {
		resp := NewResponse(tt.code, tt.text)

		if resp.Code() != tt.code {
			t.Errorf("expected code %d, got %d", tt.code, resp.Code())
		}

		if resp.Text() != tt.text {
			t.Errorf("expected text %q, got %q", tt.text, resp.Text())
		}
	}
}
