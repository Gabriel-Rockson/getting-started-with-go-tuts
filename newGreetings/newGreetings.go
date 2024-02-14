package newGreetings

import "fmt"

func Hello(fullName string) string {
	message := fmt.Sprintf("%s, this is a new greeting.", fullName)

	return message
}
