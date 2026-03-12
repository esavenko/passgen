package constants

type Screen int

const (
	Letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits         = "0123456789"
	SpecialSymbols = "!@#$%^&*()_+-=[]{}|;:',./<>?"
)

const (
	Menu Screen = iota
	PasswordGeneration
	Settings
)

const (
	Logo = "______                              \n| ___ \\                             \n| |_/ /_ _ ___ ___  __ _  ___ _ __  \n|  __/ _` / __/ __|/ _` |/ _ \\ '_ \\ \n| | | (_| \\__ \\__ \\ (_| |  __/ | | |\n\\_|  \\__,_|___/___/\\__, |\\___|_| |_|\n                    __/ |           \n                   |___/            "
)
