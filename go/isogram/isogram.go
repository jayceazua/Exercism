package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if the word input is a isogram
func IsIsogram(w string) bool {
	// memoization mapping
	runes := map[rune] bool {}
	// make all the characters in the string lowercase
	wl := strings.ToLower(w)
	for _, r := range wl {
		// check if it is a letter
		if unicode.IsLetter(r) && runes[r] {
			return false
		}
		runes[r] = true
	}
	return true
}
