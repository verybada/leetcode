import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "nums, expected",
    [
        [[1, 1, 2, 2, 2, 3], [3, 1, 1, 2, 2, 2]],
        [[2, 3, 1, 3, 2], [1, 3, 3, 2, 2]],
        [[-1, 1, -6, 4, 5, -6, 1, 4, 1], [5, -1, 4, 4, -6, -6, 1, 1, 1]],
    ],
)
def test_frequency_sort(nums, expected):
    assert Solution().frequencySort(nums) == expected
