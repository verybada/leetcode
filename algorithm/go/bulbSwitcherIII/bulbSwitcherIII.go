package bulbSwitcherIII

func numTimesAllBlue(light []int) int {
	allBlueCount := 0
	farestLightOnBlub := 0

	for index, lightOnBlub := range light {
		if lightOnBlub > farestLightOnBlub {
			farestLightOnBlub = lightOnBlub
		}

		if farestLightOnBlub == index+1 {
			allBlueCount += 1
		}
	}

	return allBlueCount
}

func numTimesAllBlueX(light []int) int {
	allBlueCount := 0
	blubsLen := len(light) + 1
	blubs := make([]int, blubsLen)

	farestLightOnBlub := 0

	checkAllBlue := func(start int, end int) bool {
		for i := start; i >= end; i-- {
			if blubs[i] == 0 {
				return false
			}
		}
		return true
	}
	for _, lightOnBlub := range light {
		blubs[lightOnBlub] = 1

		if lightOnBlub > farestLightOnBlub {
			farestLightOnBlub = lightOnBlub
			if checkAllBlue(lightOnBlub, 1) {
				allBlueCount += 1
			}
		} else if farestLightOnBlub >= lightOnBlub {
			if checkAllBlue(farestLightOnBlub, 1) {
				allBlueCount += 1
			}
		}

	}

	return allBlueCount
}
