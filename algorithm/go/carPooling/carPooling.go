package carpooling

import (
	"sort"
)

func carPooling(trips [][]int, capacity int) bool {
	records := make(map[int]int)
	for _, trip := range trips {
		passengers := trip[0]
		pickupLocation := trip[1]
		dropoffLocation := trip[2]

		value, ok := records[pickupLocation]
		if !ok {
			value = 0
		}
		records[pickupLocation] = value - passengers

		value, ok = records[dropoffLocation]
		if !ok {
			value = 0
		}
		records[dropoffLocation] = value + passengers
	}

	locations := make([]int, 0, len(records))
	for key, _ := range records {
		locations = append(locations, key)
	}
	sort.Ints(locations)
	for _, location := range locations {
		passengers := records[location]
		capacity += passengers
		if capacity < 0 {
			return false
		}
	}
	return true
}
