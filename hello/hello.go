package main

import (
	"fmt"

	"github.com/gabgrz/golang/greetings"
)

func main() {
	message := greetings.Hello("Gabriel")
	fmt.Println(message)
}
