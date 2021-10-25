package main

import "fmt"

func main() {
	//var s string    ---> declaration (initialized to empty string by compiler by default)
	//s = "Hello, World!"  ----> initialization
	s := "Hello, World!" //declare and initialize together
	fmt.Println(s)       //Outputs "Hello, World!"

	s = "Hello, \n\t\"world!\" with a backslash\\"
	fmt.Println(s)
	/* Output:
		Hello,
	        "world!" with a backslash\
	*/

	s = `Hello,
	"world!" with a backslash \`
	fmt.Println(s)
	/* Output:
		Hello,
	        "world!" with a backslash\
	*/

	//Go supports printing emojis
	s1 := "Hi there!"
	s2 := "👋🤗"
	fmt.Println(s1 + " " + s2) //outputs Hi there! 👋🤗

	s = "Hello, world!"
	b := s[0]
	b2 := s[4]
	fmt.Println(s, b, string(b), b2, string(b2)) //outputs Hello, world! 72 H 111 o --> ascii value of 'H' is 72 and 'o' is 111

	//substring of s
	sub1 := s[0:5]  //s[5] is excluded
	sub2 := s[7:12] //s[12] is excluded
	sub3 := s[:6]
	sub4 := s[7:]
	fmt.Println("s = " + s + "\n" + sub1 + "\n" + sub2 + "\n" + sub3 + "\n" + sub4)
	/*outputs:
	s = Hello, world!
	Hello
	world
	Hello,
	world!
	*/

	//length of the strings
	fmt.Println(s, len(s), sub1, len(sub1)) //outputs "Hello, world! 13 Hello 5"

	s = "👋 🤗"
	ss2 := s[:2]
	ss3 := s[:6] // 0th to 5th byte (1st four bytes is for the rune/special characters, space 1 byte and 5th byte is 1st byte of 2nd emoji)

	fmt.Println(s, len(s), "\n", ss2, len(ss2), "\n", ss3, len(ss3))
	/* outputs:
	👋 🤗 9
	�� 2
	👋 � 6

	length of s = 9 becoz each emoji takes 4 bytes and a space takes 1 byte so total 9 bytes
	�� becoz the substring is 1st 2 bytes of the string 's' which are illegal two characters in UTF-8 so two ��
	👋 � becoz 1st four bytes is for the rune/special characters which is legal in UTF-8, 1 byte for space, then 5th byte is 1st byte of 2nd emoji which is illegal in UTF-8 hence � --> Total 6 bytes so length is 6 (0th to 5th byte)
	*/

	//using rune
	s = "Hello, "
	var r rune = 127757
	s = s + string(r)
	fmt.Println(s) // outputs : Hello, 🌍

	//can also use rune like this
	s = "Hello, "
	r = '🌍'
	s = s + string(r)
	fmt.Println(s) // outputs : Hello, 🌍

	//arrays
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	fmt.Println(vals, vals[0], vals[1], vals[2]) //outputs: [2 4 6] 2 4 6

	/* Cannot assign smaller sized array into bigger sized array
	var vals2 [4]int = vals
	fmt.Println(vals, vals2) //Compilation error: cannot use vals (type [3]int) as type [4]int in assignment
	*/

}
