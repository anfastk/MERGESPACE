package valueobject

import "strings"

var commonPasswords = map[string]struct{}{
	"password":    {},
	"password123": {},
	"123456789":   {},
	"qwerty":      {},
	"letmein":     {},
	"admin":       {},
	"welcome":     {},
	"iloveyou":    {},
	"monkey":      {},
}

func isCommonPassword(p string) bool {
	_, found := commonPasswords[strings.ToLower(p)]
	return found
}
