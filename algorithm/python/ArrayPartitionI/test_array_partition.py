import pytest

from .array_partition import Solution


@pytest.mark.parametrize('nums, expected', [
    [[1, 4, 2, 3], 4],
    [[6, 2, 6, 5, 1, 2], 9],
])
def test_arrayPariSum(nums, expected):
    assert Solution().arrayPairSum(nums) == expected
