package hello

import "fmt"

func Hello(recipient, language string) string {
	greeting := "Hello"

	if recipient == "" {
		recipient = "world"
	}

	switch language {
	case "Spanish":
		greeting = "Hola"	
	}

	return fmt.Sprintf("%s, %s", greeting, recipient)
}