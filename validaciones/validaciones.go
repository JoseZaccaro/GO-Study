package validaciones

import (
	"regexp"
	"unicode"
)

var Regex_correo = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)

func ValidarPassword(s string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(s) >= 6 && len(s) <= 20 {
		hasMinLen = true
	}

	for _, r := range s {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsNumber(r):
			hasNumber = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber
}
