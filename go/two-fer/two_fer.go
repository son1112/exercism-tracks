// Package twofer is short for two for one. One for you and one for me.
package twofer

import "fmt"

// ShareWith takes a given name and returns a string using the name,
// defaulting to "you" if given string is empty.
// If the given name is "Alice", the result should be "One for Alice, one for me."
// If no name is given, the result should be "One for you, one for me."
func ShareWith(name string) string {
	message := "One for %s, one for me."

	return fmt.Sprintf(message, Addressed(name))
}

// Addressed takes a string and returns "you" if empty,
// otherwise returns given string
func Addressed(name string) string {
	switch name {
	case "":
		return "you"
	default:
		return name
	}
}
