package config

import "testing"

func TestNewDefaultSettings(t *testing.T) {
	s := NewDefaultSettings()

	if s == nil {
		t.Fatal("NewDefaultSettings() returned nil")
	}

	tests := []struct {
		name string
		got  any
		want any
	}{
		{"Length", s.Length, 15},
		{"UseDigits", s.UseDigits, true},
		{"UseSpecialSymbols", s.UseSpecialSymbols, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}
