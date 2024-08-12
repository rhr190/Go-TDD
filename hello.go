package main

import "fmt"

const englishHelooPrefix = "Hello, "

func Hello(recipient string) string {
	if recipient == "" {
		recipient = "world"
	}
	return englishHelooPrefix + recipient
}

func main() {
	fmt.Println(Hello(""))
}
