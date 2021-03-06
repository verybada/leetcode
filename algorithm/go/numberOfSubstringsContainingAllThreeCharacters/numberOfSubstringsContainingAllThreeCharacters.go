package numberofsubstringscontainingallthreecharacters

func numberOfSubstrings(s string) int {
	chars := []rune(s)
	length := len(chars)
	subStringCount := 0
	charCount := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
	}

	left := 0
	for right, char := range chars {
		charCount[char]++
		for containABC(charCount) {
			subStringCount += length - right

			charCount[chars[left]]--
			left++
		}
	}
	return subStringCount
}

func containABC(charCount map[rune]int) bool {
	for _, count := range charCount {
		if count == 0 {
			return false
		}
	}
	return true
}

func numberOfSubstringsBruteForce(s string) int {
	chars := []rune(s)
	subStringCount := 0
	for size := len(chars); size >= 3; size-- {
		subStringCount += calculateMatchedSubstring(chars, size)
	}

	return subStringCount
}

func calculateMatchedSubstring(chars []rune, subStringSize int) int {
	window := NewRuneWindow()
	subStringCount := 0
	left := 0
	right := 0
	for ; right < subStringSize; right++ {
		char := chars[right]
		window.Add(char)
	}

	for {
		if window.ContainsABC() {
			subStringCount++
		}

		if right > len(chars)-1 {
			break
		}
		rightChar := chars[right]
		window.Add(rightChar)
		leftChar := chars[left]
		window.Remove(leftChar)
		right++
		left++
	}
	return subStringCount
}

func NewRuneWindow() *runeWindow {
	return &runeWindow{
		runeCounts: map[rune]int{
			'a': 0,
			'b': 0,
			'c': 0,
		},
	}
}

type runeWindow struct {
	runeCounts map[rune]int
}

func (r *runeWindow) Add(char rune) {
	r.runeCounts[char]++
}

func (r *runeWindow) Remove(char rune) {
	r.runeCounts[char]--
}

func (r *runeWindow) ContainsABC() bool {
	for _, count := range r.runeCounts {
		if count == 0 {
			return false
		}
	}
	return true
}
