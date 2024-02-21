package hello

import (
	"fmt"

	"example.com/greetings"
)

func SayHelloMadam() {
	message, err := greetings.Hello("Ma'am")

	if err == nil {
		fmt.Println(message)
	}
}

func SayHello(name string) {

	message, err := greetings.Hello(name)

	if err != nil {
		fmt.Printf("Error occured when using name: %s \n", name)

		fmt.Println(err)
	}

	fmt.Println(message)
	return
}
