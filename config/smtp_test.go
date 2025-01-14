package config

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {

	tests := []struct {
		name          string
		config        SmtpConfig
		expectedError error
	}{
		{
			name: "empty host",
			config: SmtpConfig{
				Host: "",
			},
			expectedError: fmt.Errorf("smtp host is required"),
		},
		{
			name: "valid config",
			config: SmtpConfig{
				Host: "localhost",
			},
			expectedError: nil,
		},
		{
			name: "valid config with Ip",
			config: SmtpConfig{
				Ip: "127.0.0.1",
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.config.validate()
			if test.expectedError == nil && err == nil {
				return
			}
			if errors.Is(err, test.expectedError) {
				t.Errorf("expected error %v, got %v", test.expectedError, err)
			}
		})
	}
}
