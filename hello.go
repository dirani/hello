package main

import (
	"fmt"

	"github.com/dirani/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("Gladys")
	fmt.Println(message)
	if err != nil {
		fmt.Println(greetings.Hello(""))
	} else {
		fmt.Println(err)
	}
}
