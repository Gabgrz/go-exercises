package main

import (
	"fmt"
	"log"

	"github.com/Gabgrz/golang/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a correct greeting message
	goodMessage, err := greetings.Hello("Gabriel")
	fmt.Println(goodMessage)

	// Request a Hello with randomFormat
	randomMessage, err := greetings.HelloRandom("Alejandro")
	fmt.Println(randomMessage)

	// Request an incorrect greeting message
	badMessage, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(badMessage)

}
