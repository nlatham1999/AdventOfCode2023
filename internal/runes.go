package internal

func RuneToAsciiNumber(r rune) int {
	switch {
	case r >= 'a' && r <= 'z':
		// 'a' - 'a' + 1 = 1, 'b' - 'a' + 1 = 2, ..., 'z' - 'a' + 1 = 26
		return int(r - 'a' + 1)
	case r >= 'A' && r <= 'Z':
		// 'A' - 'A' + 27 = 27, 'B' - 'A' + 27 = 28, ..., 'Z' - 'A' + 27 = 52
		return int(r - 'A' + 27)
	default:
		// Returns 0 if the rune is not in a-z or A-Z
		return 0
	}
}
