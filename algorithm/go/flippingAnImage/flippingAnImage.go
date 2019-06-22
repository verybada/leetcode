package flippingAnImage

import (
)

func flipAndInvertImage(A [][]int) [][]int {
    for _, bitmap := range A {
        bitmapLen := len(bitmap)

        midIndex := bitmapLen / 2
        if bitmapLen % 2 == 0 {
            midIndex = (bitmapLen / 2) - 1
        }

        for i:=0; i<=midIndex; i++ {
            tail := bitmapLen-i-1
            if i != tail {
                bitmap[tail], bitmap[i] = bitmap[i]^1, bitmap[tail]^1
            } else {
                bitmap[i] = bitmap[i]^1
            }
        }
    }
    return A
}
