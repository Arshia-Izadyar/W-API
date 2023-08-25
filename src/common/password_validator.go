package common

import (
	"unicode"
	"wapi/src/config"
)

func HasLetter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, c := range s {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, c := range s {
		if unicode.IsLower(c) && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func HasUpper(s string) bool {
	for _, c := range s {
		if unicode.IsUpper(c) && unicode.IsLetter(c) {
			return true
		}
	}
	return false
}

func CheckPassword(pass string) bool {
	cfg := config.LoadCfg()
	if len(pass) < cfg.PassWord.MinLength || len(pass) > cfg.PassWord.MaxLength {
		return false
	}
	if cfg.PassWord.IncludeChars && !HasLetter(pass) {
		return false
	}
	if cfg.PassWord.IncludeDigits && !HasDigits(pass) {
		return false
	}
	if cfg.PassWord.IncludeUppercase && !HasUpper(pass) {
		return false
	}
	if cfg.PassWord.IncludeLowercase && !HasLower(pass) {
		return false
	}

	return true
}
