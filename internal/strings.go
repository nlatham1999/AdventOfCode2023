package internal

func SplitString(str string) (first string, second string) {
	runes := []rune(str)

	// Calculate the middle index
	mid := len(runes) / 2

	// Split the rune slice into two halves
	firstHalf := runes[:mid]
	secondHalf := runes[mid:]

	// Convert rune slices back to strings
	firstHalfStr := string(firstHalf)
	secondHalfStr := string(secondHalf)

	return firstHalfStr, secondHalfStr
}

func StringToCharacters(str string) []rune {
	return []rune(str)
}
