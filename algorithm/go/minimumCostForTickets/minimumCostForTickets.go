package mininumCostForTickets

func mincostTickets(days []int, costs []int) int {
	// NOTE: recursive solution, but too slow
	// return estimateMinCost(days, costs, 0, 0, 0)

	// NOTE: dynamic programming solution
	oneDayPassCost := costs[0]
	weekPassCost := costs[1]
	monthPassCost := costs[2]

	dayMap := map[int]bool{}
	for _, day := range days {
		dayMap[day] = true
	}

	// shift one day, make index start from 1
	dayCost := make([]int, 366)
	for i := 1; i < 366; i++ {
		_, ok := dayMap[i]
		if !ok {
			// not a travel day, cost like yesterday
			dayCost[i] = dayCost[i-1]
			continue
		}

		// when use one day pass, cost is yesterdayCost add one day pass cost
		minCost := dayCost[i-1] + oneDayPassCost

		// when use 7 day pass, cost is weekAgoCost add 7 day pass cost
		// if 7 day ago not exist, assume the cost is zero
		weekAgoCost := 0
		if i-7 > 0 {
			weekAgoCost = dayCost[i-7]
		}
		cost := weekAgoCost + weekPassCost
		if cost < minCost {
			minCost = cost
		}

		// when use 30 day pass, cost is monthAgoCost add 30 day pass cost
		// if 30 day ago not exist, assume the cost is zero
		monthAgoCost := 0
		if i-30 > 0 {
			monthAgoCost = dayCost[i-30]
		}
		cost = monthAgoCost + monthPassCost
		if cost < minCost {
			minCost = cost
		}

		dayCost[i] = minCost

	}

	return dayCost[365]
}

func estimateMinCost(
	days []int, costs []int, index int, currentCost int, validUntil int,
) int {
	if index >= len(days) {
		return currentCost
	}

	day := days[index]
	nextIndex := index + 1

	if validUntil >= day {
		return estimateMinCost(days, costs, nextIndex, currentCost, validUntil)
	}

	oneDayPassCost := costs[0]
	weekPassCost := costs[1]
	monthPassCost := costs[2]

	minCost := estimateMinCost(
		days, costs, nextIndex, currentCost+oneDayPassCost, day)
	cost := estimateMinCost(
		days, costs, nextIndex, currentCost+weekPassCost, day+6)
	if cost < minCost {
		minCost = cost
	}

	cost = estimateMinCost(
		days, costs, nextIndex, currentCost+monthPassCost, day+29)
	if cost < minCost {
		minCost = cost
	}

	return minCost
}
