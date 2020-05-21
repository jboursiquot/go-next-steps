package stringutils

import "strings"

// Caser handles uppercasing and lowercasing strings.
type Caser struct{}

// Upper returns uppercased string.
func (Caser) Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower returns lowercase string.
func (Caser) Lower(s string) string {
	return strings.ToLower(s)
}
