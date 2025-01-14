package config

import (
	"os"
	"path/filepath"
	"testing"
)

var sampleConfig = `
[smtp]
host = "smtp.example.com"
`

func TestInit(t *testing.T) {
	// Create sample file
	configDir := filepath.Join(t.TempDir(), "config.toml")
	f, err := os.Create(configDir)
	if err != nil {
		t.Errorf("Error creating file: %s", err)
	}
	f.Write([]byte(sampleConfig))

	c, err := Init(configDir)

	if err != nil {
		t.Errorf("Error reading config file: %s", err)
	}

	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"host", c.Smtp.Host, "smtp.example.com"},
	}
	for _, tt := range tests {
		if tt.value != tt.expected {
			t.Errorf("Expected %s, got %s", tt.expected, tt.value)
		}
	}
}
