package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

func rotate13(in rune) rune {
	if in >= 'A' && in <= 'Z' {
		return ((((in - 'A') + 13) % 26) + 'A')
	}
	if in >= 'a' && in <= 'z' {
		return ((((in - 'a') + 13) % 26) + 'a')
	}
	return in
}

/* strings.Map() Function in Golang is used to returns a copy of the string given string with all its characters modified according to the mapping function. If mapping returns a negative value, the character is dropped from the string with no replacement.
Syntax:

func Map(mapping func(rune) rune, s string) string
*/

func main() {
	//Ex1: Rotating a string to 13 characters
	s := "This is Rachana 😃"
	s2 := strings.Map(rotate13, s)
	fmt.Println(s2)                 //outputs: Guvf vf Enpunan 😃
	s3 := strings.Map(rotate13, s2) //outputs: This is Rachana 😃
	fmt.Println(s3)

	//Ex2: Replacing all 'e' with @ in a given string
	modified := func(r rune) rune {
		if r == 'e' {
			return '@'
		}
		return r
	}
	input := "Hello, Welcome to GeeksforGeeks"
	fmt.Println("Input string: " + input) //outputs: Input string: Hello, Welcome to GeeksforGeeks
	// using the function
	result := strings.Map(modified, input)
	fmt.Println("Output string: " + result) //outputs: Output string: H@llo, W@lcom@ to G@@ksforG@@ks

	//Ex3: use utf8.RuneCountInString to know the actual count of characters in a string
	s1 := "👋 🌏"
	fmt.Println(s1)                         //outputs: 👋 🌏
	fmt.Println(len(s1))                    //outputs: 9  (4+1+4, 4 bytes for each rune and 1 byte for a space)
	fmt.Println(utf8.RuneCountInString(s1)) //outputs: 3

	//Ex4: use of time.Now package
	fmt.Println(time.Now()) //outputs: 2021-11-09 12:43:46.8996721 +0530 IST m=+0.004966301
	nanos := time.Now().UnixNano()
	fmt.Println(nanos) //outputs: 1636442026901690100

	//Ex5: use of math and math/rand packages
	rand.Seed(time.Now().UnixNano()) //initializing 	rand.Seed so that each time program is run,different random number is selected
	a := rand.Intn(10)
	b := rand.Intn(10)
	c := int(math.Max(float64(a), float64(b)))
	fmt.Println(a, b)           //outputs: 4 8 in 1st run, 0 9 in 2nd run
	fmt.Println(c, "is bigger") //outputs: "8 is bigger" in 1st run, "9 is bigger" in 2nd run, each time program is run, different random number is selected.

}
