package generator

import (
	"math"
	"strings"
	"testing"

	"github.com/esavenko/passgen/internal/constants"
)

func TestPoolSize(t *testing.T) {
	tests := []struct {
		name string
		cfg  GeneratorConfig
		want int
	}{
		{"letters only", GeneratorConfig{Length: 10, UseDigits: false, UseSpecialSymbols: false}, len(constants.Letters)},
		{"letters + digits", GeneratorConfig{Length: 10, UseDigits: true, UseSpecialSymbols: false}, len(constants.Letters) + len(constants.Digits)},
		{"letters + special", GeneratorConfig{Length: 10, UseDigits: false, UseSpecialSymbols: true}, len(constants.Letters) + len(constants.SpecialSymbols)},
		{"all", GeneratorConfig{Length: 10, UseDigits: true, UseSpecialSymbols: true}, len(constants.Letters) + len(constants.Digits) + len(constants.SpecialSymbols)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PoolSize(tt.cfg); got != tt.want {
				t.Errorf("PoolSize() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestEntropy(t *testing.T) {
	tests := []struct {
		name string
		cfg  GeneratorConfig
		want float64
	}{
		{"zero length", GeneratorConfig{Length: 0, UseDigits: true, UseSpecialSymbols: true}, 0},
		{"negative length", GeneratorConfig{Length: -1, UseDigits: true, UseSpecialSymbols: true}, 0},
		{"letters only len 10", GeneratorConfig{Length: 10, UseDigits: false, UseSpecialSymbols: false}, 10 * math.Log2(float64(len(constants.Letters)))},
		{"all chars len 15", GeneratorConfig{Length: 15, UseDigits: true, UseSpecialSymbols: true}, 15 * math.Log2(float64(len(constants.Letters)+len(constants.Digits)+len(constants.SpecialSymbols)))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Entropy(tt.cfg)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("Entropy() = %f, want %f", got, tt.want)
			}
		})
	}
}

func TestStrength(t *testing.T) {
	tests := []struct {
		name    string
		entropy float64
		want    string
	}{
		{"very weak low", 0, "Very Weak"},
		{"very weak high", 27.9, "Very Weak"},
		{"weak low", 28, "Weak"},
		{"weak high", 35.9, "Weak"},
		{"reasonable low", 36, "Reasonable"},
		{"reasonable high", 59.9, "Reasonable"},
		{"strong low", 60, "Strong"},
		{"strong high", 127.9, "Strong"},
		{"very strong", 128, "Very Strong"},
		{"very strong high", 256, "Very Strong"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strength(tt.entropy); got != tt.want {
				t.Errorf("Strength(%f) = %q, want %q", tt.entropy, got, tt.want)
			}
		})
	}
}

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name    string
		cfg     GeneratorConfig
		allowed string
	}{
		{"letters only", GeneratorConfig{Length: 20, UseDigits: false, UseSpecialSymbols: false}, constants.Letters},
		{"letters + digits", GeneratorConfig{Length: 20, UseDigits: true, UseSpecialSymbols: false}, constants.Letters + constants.Digits},
		{"letters + special", GeneratorConfig{Length: 20, UseDigits: false, UseSpecialSymbols: true}, constants.Letters + constants.SpecialSymbols},
		{"all", GeneratorConfig{Length: 20, UseDigits: true, UseSpecialSymbols: true}, constants.Letters + constants.Digits + constants.SpecialSymbols},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := GeneratePassword(tt.cfg)
			if err != nil {
				t.Fatalf("GeneratePassword() error: %v", err)
			}

			if len(password) != tt.cfg.Length {
				t.Errorf("password length = %d, want %d", len(password), tt.cfg.Length)
			}

			for _, ch := range password {
				if !strings.ContainsRune(tt.allowed, ch) {
					t.Errorf("password contains %q which is not in allowed set", ch)
				}
			}
		})
	}
}

func TestGeneratePasswordUniqueness(t *testing.T) {
	cfg := GeneratorConfig{Length: 30, UseDigits: true, UseSpecialSymbols: true}

	p1, _ := GeneratePassword(cfg)
	p2, _ := GeneratePassword(cfg)

	if p1 == p2 {
		t.Error("two generated passwords are identical — likely not random")
	}
}
