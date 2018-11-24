// Package proverb produces an array of proverbs from an array of inputs
package proverb

// Proverb outputs an array of proverbial rhymes from given array of inputs
func Proverb(rhyme []string) []string {
	length := len(rhyme)
	rhymes := []string{}

	switch length {
	case 0:
		return []string{}
	case 1:
		"And all for the want of a "
	}

	for index, item := range rhyme {
		// multiple strings, all but last like
		// "For want of a <item> the <next-item> was lost."
		// last string like "And all for the want of a <first-item>."
		// no input defaults to empty string?
	}
}
