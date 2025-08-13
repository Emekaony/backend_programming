package utils

import (
	"strings"
	"unicode/utf8"
)

/*

	remember that string(integer) does not convert the integer to the
	string representation of that integer. Instead, it produces a character that mathces
	the unicode code point of the integer

	if u want to convert from integer to string, use strconv.Itoa(number)
*/

// / returns true if first ends with second in a case-insensitive manner
func EndsWith(first, second string) bool {
	diff := utf8.RuneCountInString(first) - utf8.RuneCountInString(second)
	return strings.EqualFold(first[diff:], second)
}
