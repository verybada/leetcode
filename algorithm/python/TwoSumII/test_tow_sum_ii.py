import pytest

from .two_sum_ii import Solution


@pytest.mark.parametrize('numbers, target, expected', [
    [[2, 7, 11, 15], 9, [1, 2]],
    [[2, 3, 4], 6, [1, 3]],
    [[-1, 0], -1, [1, 2]]

])
def test_twoSum(numbers, target, expected):
    result = Solution().twoSum(numbers, target)
    assert result == expected
