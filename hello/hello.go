package main

import (
	"fmt"

	"github.com/Gabgrz/golang/greetings"
)

func main() {
	message := greetings.Hello("Gabriel")
	fmt.Println(message)
}
