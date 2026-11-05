package inttoroman

func intToRoman(input int) string {
	romanString := ""
	mapper := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, value := range values {
		roman := mapper[value]
		for value <= input {
			input -= value
			romanString += roman

			if input == 0 {
				return romanString
			}
		}
	}
	return ""
}
