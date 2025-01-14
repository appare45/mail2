package config

import "testing"

func TestSmtpDefaultConfig(t *testing.T) {
	config := Config{}
	config.Smtp.defaultConfig()
	if config.Smtp.Port != 587 {
		t.Errorf("Expected 587, got %d", config.Smtp.Port)
	}
}
