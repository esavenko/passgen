package generator

import (
	"crypto/rand"
	"math"
	"math/big"
	"strings"

	"github.com/esavenko/passgen/internal/constants"
)

type GeneratorConfig struct {
	Length            int
	UseDigits         bool
	UseSpecialSymbols bool
}

func GeneratePassword(cfg GeneratorConfig) (string, error) {
	var charPool string

	charPool = constants.Letters

	if cfg.UseDigits {
		charPool += constants.Digits
	}

	if cfg.UseSpecialSymbols {
		charPool += constants.SpecialSymbols
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

func PoolSize(cfg GeneratorConfig) int {
	size := len(constants.Letters)
	if cfg.UseDigits {
		size += len(constants.Digits)
	}
	if cfg.UseSpecialSymbols {
		size += len(constants.SpecialSymbols)
	}
	return size
}

func Entropy(cfg GeneratorConfig) float64 {
	poolSize := PoolSize(cfg)
	if poolSize <= 1 || cfg.Length <= 0 {
		return 0
	}
	return float64(cfg.Length) * math.Log2(float64(poolSize))
}

func Strength(entropy float64) string {
	switch {
	case entropy < 28:
		return "Very Weak"
	case entropy < 36:
		return "Weak"
	case entropy < 60:
		return "Reasonable"
	case entropy < 128:
		return "Strong"
	default:
		return "Very Strong"
	}
}
