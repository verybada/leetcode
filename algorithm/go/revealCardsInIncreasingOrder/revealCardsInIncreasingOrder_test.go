package main

import (
	"reflect"
	"testing"
)

func TestRevealCardsInIncresingOrder(t *testing.T) {
	input := []int{17, 13, 11, 2, 3, 5, 7}
	output := []int{2, 13, 3, 11, 5, 17, 7}

	result := deckRevealedIncreasing(input)
	if !reflect.DeepEqual(result, output) {
		t.Fatalf("expected %v but %v", output, result)
	}
}
