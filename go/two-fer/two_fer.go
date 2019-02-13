package twofer

import "fmt"

// === Credit from: https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing
// === Credit from: https://www.golang-book.com/books/intro/4
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
