package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//Get a greetings message and print them
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darring"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
