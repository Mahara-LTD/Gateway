package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Ma'am")
	fmt.Println(message)
}
