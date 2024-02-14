package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("%s You are welcome to our little part of the world", name)

	return message
}
