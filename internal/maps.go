package internal

func GetOccurencesOfStringArray(occurences map[string]int, arr []string) {
	for _, x := range arr {
		val, found := occurences[x]
		if found {
			occurences[x] = val + 1
		} else {
			occurences[x] = 0
		}
	}
}

func GetOccurencesOfRuneArray(occurences map[rune]int, arr []rune) {
	for _, x := range arr {
		val, found := occurences[x]
		if found {
			occurences[x] = val + 1
		} else {
			occurences[x] = 0
		}
	}
}

func GetOccurencesOfString(occurences map[rune]int, str string) {
	runes := StringToCharacters(str)
	GetOccurencesOfRuneArray(occurences, runes)
}

func StringsToSet(str []string) map[string]interface{} {
	set := make(map[string]interface{})
	for _, x := range str {
		set[x] = nil
	}
	return set
}

func SingleStringToSet(str string) map[rune]interface{} {
	r := []rune(str)
	return RunesToSet(r)
}

func IntToSet(ints []int) map[int]interface{} {
	set := make(map[int]interface{})
	for _, x := range ints {
		set[x] = nil
	}
	return set
}

func RunesToSet(runes []rune) map[rune]interface{} {
	set := make(map[rune]interface{})
	for _, x := range runes {
		set[x] = nil
	}
	return set
}
