package longpressedname

func isLongPressedName(name string, typed string) bool {
	nIndex := 0
	nMax := len(name)
	tIndex := 0
	tMax := len(typed)

	// typed less than name
	if tMax < nMax {
		return false
	}

	for nIndex < nMax {
		n := name[nIndex]

		// name not all matched but typed is gone
		if tIndex >= tMax {
			return false
		}

		t := typed[tIndex]
		if n == t {
			nIndex++
			tIndex++
			continue
		}

		// if not equal, move back to remove redundant char
		previousIndex := nIndex - 1
		// mismatch for first char
		if previousIndex < 0 {
			return false
		}
		previousN := name[previousIndex]
		// not redundant, is totally different
		if previousN != t {
			return false
		}

		for previousN == t {
			tIndex++
			// try to remove redundant but out of index
			if tIndex >= tMax {
				return false
			}
			t = typed[tIndex]
		}
	}

	// when name are matched, needs to make sure typed are repeat
	// last char of N
	lastN := name[nMax-1]
	for ; tIndex < tMax; tIndex++ {
		t := typed[tIndex]
		if lastN != t {
			return false
		}
	}
	return true
}
