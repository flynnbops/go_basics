package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// Functions starting with an Uppercase letter can be called from functions not in the same package.
// This is called the "exported name"
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}