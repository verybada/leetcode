package main

import ()

func judgeCircle(moves string) bool {
	// R L U D
	var horizon int
	var vertical int
	for _, char := range moves {
		switch char {
		case 'R':
			vertical += 1
			break
		case 'L':
			vertical -= 1
			break
		case 'U':
			horizon += 1
			break
		case 'D':
			horizon -= 1
			break
		}
	}

	return horizon == 0 && vertical == 0
}
