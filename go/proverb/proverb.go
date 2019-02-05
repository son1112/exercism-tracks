// Package proverb produces an array of proverbs from an array of inputs
package proverb

import (
	"fmt"
)

// Proverb outputs an array of proverbial rhymes from given array of inputs
func Proverb(rhyme []string) []string {
	length := len(rhyme)
	rhymes := make([]string, length)

	fmt.Printf("LENGTH: %v \n", length)

	// if length == 0 {
	// 	return rhymes
	// }

	// for index, item := range rhyme {
	// 	statement := ""
	// 	switch index {
	// 	case 0:
	// 		statement = fmt.Sprintf("And all for the want of a %s.", item)
	// 	default:
	// 		statement = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index-1], item)
	// 	}

	// 	rhymes = append([]string{statement}, rhymes...)
	switch length {
	case 0:
		rhymes = []string{}
	case 1:
		rhymes = []string{fmt.Sprintf("And all for the want of a %v.", rhyme[0])}
	default:
		collect := rhyme[:len(rhyme)-1]

		fmt.Printf("COLLECT: %v \n", collect)

		for index, item := range collect {
			fmt.Printf("ITEM: %v \n", item)
			fmt.Printf("INDEX: %v \n", index)
			fmt.Printf("LENGTH: %v \n", len(collect))

			if index+1 <= len(collect)-1 {
				next_item := collect[index+1]
				parable := fmt.Sprintf("For want of a %v the %v was lost.", item, next_item)
				rhymes = append(rhymes, parable)

				fmt.Printf("ITEM: %v", item)
			}
		}

		first := rhyme[0]
		string := fmt.Sprintf("And all for the want of a %v", first)
		rhymes = append(rhymes, string)
	}

	return rhymes
}
