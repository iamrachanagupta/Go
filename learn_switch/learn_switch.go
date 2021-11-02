package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	if word == "Hello" {
		fmt.Println("Hi there")
	} else if word == "guten tag" {
		fmt.Println("Danke")
	} else if word == "gute nacht" {
		fmt.Println("Good night to u too") //Outputs this if passed "gute nacht" , quotes necessary for two or more words while passing arguments
	} else {
		fmt.Println("Sorry, I didnt get you!")
	}

	//Now same thing using a switch statements
	greet := "greetings"
	switch word {
	case "Hi":
		fmt.Println("Hello!")
		fallthrough /* this keyword means it will continue with the next case statement action as well Because of which if passed "Hi" then output will be :
		Hello!
		Hi there
		*/
	case "Hello":
		fmt.Println("Hi there")
	case "guten tag", "good day": //either passed "guten tag" or "good day", output will be "Danke"
		fmt.Println("Danke")
	case "gute nacht":
		fmt.Println("Good night to u too") //Outputs "Good night to u too" if passed "gute nacht" , quotes necessary for two or more words while passing arguments
	case "farewell": // since no statement proceed this case statement, if farewell is passed it doesnt give any output.
	case greet: //passing greet variable.
		fmt.Println("Salutations")
	default:
		fmt.Println("Sorry, I didnt get you!")
	}

	//switch with some initialization, that is accessible only within switch
	switch l := len(word); word {
	case "Hi":
		fmt.Println("Hello!")
		fallthrough
	case "Hello":
		fmt.Println("Hi there")
	case "guten tag", "good day":
		fmt.Println("Danke")
	case "gute nacht":
		fmt.Println("Good night to u too")
	case "farewell":
	default:
		fmt.Println("Sorry, I didnt get you!, But it was ", l, "characters long")
		/*
			command: go run .\learn_switch.go "Wei gehts"
			Outputs:
			Sorry, I didnt get you!, But it was  9 characters long
		*/
	}
}
