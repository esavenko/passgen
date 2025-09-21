package app

import (
	"crypto/rand"
	"math/big"
	"strings"
)

type GeneratorConfig struct {
	Length            int
	UseDigits         bool
	UseSpecialSymbols bool
}

func GeneratePassword(cfg GeneratorConfig) (string, error) {
	var charPool string

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	symbols := "!@#$%^&*()_+-=[]{}|;:',./<>?"

	charPool = letters

	if cfg.UseDigits {
		charPool += digits
	}

	if cfg.UseSpecialSymbols {
		charPool += symbols
	}

	var password strings.Builder
	poolLength := big.NewInt(int64(len(charPool)))

	for i := 0; i < cfg.Length; i++ {
		idx, err := rand.Int(rand.Reader, poolLength)
		if err != nil {
			return "", err
		}

		password.WriteByte(charPool[idx.Int64()])
	}

	return password.String(), nil
}
