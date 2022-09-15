package decodeMessage

func decodeMessage(key string, message string) string {
	plaintext := make([]rune, 0)
	keyMap := make(map[rune]rune)
	lastRune := 'a'
	for _, char := range key {
		if char == ' ' {
			continue
		}
		_, ok := keyMap[char]
		if !ok {
			keyMap[char] = lastRune
			lastRune += 1
		}

		if len(keyMap) == 26 {
			break
		}
	}

	keyMap[' '] = ' '
	for _, char := range message {
		plain := keyMap[char]
		plaintext = append(plaintext, plain)
	}

	return string(plaintext)
}
