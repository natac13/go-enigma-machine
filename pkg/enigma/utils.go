package enigma

func runeToAlphabetIndex(r rune) int {
	return int(r - 'A')
}

func alphabetIndexToRune(i int) rune {
	return rune(i + 'A')
}
