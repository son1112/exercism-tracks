// Package proverb produces an array of proverbs from an array of inputs
package proverb

import (
	"fmt"
)

// Proverb outputs an array of proverbial rhymes from given array of inputs
func Proverb(rhyme []string) []string {
	length := len(rhyme)
	rhymes := []string{}

	if length == 0 {
		return rhymes
	}

	for index, item := range rhyme {
		statement := ""
		switch index {
		case 0:
			statement = fmt.Sprintf("And all for the want of a %s.", item)
		default:
			statement = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index-1], item)
		}

		rhymes = append([]string{statement}, rhymes...)
	}

	return rhymes
}
