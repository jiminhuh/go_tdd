package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHellowPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if language == french {
	// 	return frenchHellowPrefix + name
	// }

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHellowPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}
