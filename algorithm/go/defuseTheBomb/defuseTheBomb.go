package defusethebomb

func decrypt(code []int, k int) []int {
	circularCode := NewCircularCode(code, 0)
	password := make([]int, len(code))
	for i := 0; i < len(code); i++ {
		circularCode.SetIndex(i)
		if k == 0 {
			password[i] = 0
		} else if k > 0 {
			password[i] = SumNextOfN(circularCode, k)
		} else if k < 0 {
			password[i] = SumPreviousOfN(circularCode, -k)
		}
	}
	return password
}

func SumNextOfN(circularCode *circularCode, nextOfN int) int {
	sum := 0
	for nextOfN > 0 {
		sum += circularCode.Next()
		nextOfN--
	}
	return sum
}

func SumPreviousOfN(circularCode *circularCode, previousOfN int) int {
	sum := 0
	for previousOfN > 0 {
		sum += circularCode.Previous()
		previousOfN--
	}
	return sum
}

func NewCircularCode(code []int, currentIndex int) *circularCode {
	return &circularCode{
		code, len(code) - 1, currentIndex,
	}
}

type circularCode struct {
	code         []int
	maxIndex     int
	currentIndex int
}

func (c *circularCode) Previous() int {
	c.currentIndex--
	if c.currentIndex < 0 {
		c.currentIndex = c.maxIndex
	}
	return c.code[c.currentIndex]
}

func (c *circularCode) Next() int {
	c.currentIndex++
	if c.currentIndex > c.maxIndex {
		c.currentIndex = 0
	}
	return c.code[c.currentIndex]
}

func (c *circularCode) SetIndex(index int) {
	c.currentIndex = index
}
