package elementAppearingMoreThanOneQuarterInSortedArray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindSpecialInteger(t *testing.T) {
	require.Equal(t, 6, findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}
