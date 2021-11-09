/*To create your own package, Create a new folder with the name of the package you wish to create.
In the folder created, create your go file that holds the Go package code you wish to create.It is recommended that you name your file the same name as your package name*/
package leftpad

import (
	"strings"
	"unicode/utf8"
)

//variable declared at package level. i.e the variable is visible anywhere within the package. Note: we cannot use ':=' operator to declare them. This variable starts with small letter so it makes it private to the package
var default_char = ' '

//Best practise is to use single line comment in Go for defining functions.
//Format takes in a string and an int and returns the string
//left-padded with spaces to the length of the int.
//If the string is already longer than the specified length, the original string is returned.
//Note that this function starts with a capital letter making them public i.e it can be accessed from anywhere.
func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

//FormatRune takes in a string and an int and a rune and returns the string
//left-padded with spaces to the length of the int.
//If the string is already longer than the specified length, the original string is returned.
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	} else {
		out := strings.Repeat(string(r), size-l) + s
		return out
	}
}
