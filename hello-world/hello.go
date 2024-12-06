package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(recipient, language string) string {
	if recipient == "" {
		recipient = "World"
	}
	return generatePrefix(language) + recipient
}

func generatePrefix(lang string) (pref string) {
	switch lang {
	case french:
		pref = frenchHelloPrefix
	case spanish:
		pref = spanishHelloPrefix
	default:
		pref = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Elodie", "Spanish"))
}