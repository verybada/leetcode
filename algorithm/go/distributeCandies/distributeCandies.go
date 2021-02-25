package distributecandies

func distributeCandies(candies int, numPeople int) []int {
	output := make([]int, numPeople)
	count := 1
	for candies > 0 {
		index := (count - 1) % numPeople
		if candies <= count {
			output[index] += candies
			break
		}
		output[index] += count
		candies -= count
		count++
	}
	return output
}

func distributeCandiesGoroutine(candies int, numPeople int) []int {
	candyPipe := make(chan int)
	go splitCandies(candies, candyPipe)
	output := make([]int, numPeople)
	index := 0
	for {
		count, more := <-candyPipe
		if !more {
			return output
		}
		output[index] += count
		index = (index + 1) % numPeople
	}
}

func splitCandies(candies int, pipe chan int) {
	count := 1
	for candies > 0 {
		if candies <= count {
			pipe <- candies
			close(pipe)
			return
		}
		pipe <- count
		candies -= count
		count++
	}
}
