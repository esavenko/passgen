package config

type Settings struct {
	Length            int
	UseDigits         bool
	UseSpecialSymbols bool
}

func NewDefaultSettings() *Settings {
	return &Settings{
		Length:            15,
		UseDigits:         true,
		UseSpecialSymbols: true,
	}
}
