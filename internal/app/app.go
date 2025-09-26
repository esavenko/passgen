package app

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/esavenko/passgen/common/constants"
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
