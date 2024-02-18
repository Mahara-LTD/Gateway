package hello

import (
	"fmt"

	"example.com/greetings"
)

func SayHelloMadam() {
	message := greetings.Hello("Ma'am")
	fmt.Println(message)
}
