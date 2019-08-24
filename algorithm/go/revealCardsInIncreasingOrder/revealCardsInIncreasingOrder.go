package main

import (
	"sort"
)

/*
do the reverse step to figure out correct list
input 1,2,3,4,5
output 1,5,2,4,3

reverse input
5,4,3,2,1
put value into orderedDeck and move last value to head

5
4,5
5,4
3,5,4
4,3,5
2,4,3,5
5,2,4,3
1,5,2,4,3
*/
func deckRevealedIncreasing(deck []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(deck)))
	deckLen := len(deck)
	orderedDeck := make([]int, 0, deckLen)
	for index, value := range deck {
		orderedDeck = append(orderedDeck, value)
		if index+1 == deckLen {
			break
		}
		orderedDeck = append(orderedDeck, orderedDeck[0])
		orderedDeck = orderedDeck[1:]
	}

	head := 0
	tail := deckLen - 1
	for {
		if tail <= head {
			break
		}

		orderedDeck[head], orderedDeck[tail] = orderedDeck[tail], orderedDeck[head]
		head++
		tail--
	}
	return orderedDeck
}
