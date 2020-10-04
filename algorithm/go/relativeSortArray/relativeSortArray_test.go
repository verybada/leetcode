package relativeSortArray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRelativeSortArray(t *testing.T) {
	array1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	array2 := []int{2, 1, 4, 3, 9, 6}
	expected := []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}
	require.Equal(t, expected, relativeSortArray(array1, array2))
}
