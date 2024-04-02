package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	helloPrefix := englishHelloPrefix

	switch language {
	case "French":
		helloPrefix = frenchHelloPrefix
	case "Spanish":
		helloPrefix = spanishHelloPrefix
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Bob", "French"))
}
