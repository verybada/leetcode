import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "arr, k, threshold, expected",
    [
        [[2, 2, 2, 2, 5, 5, 5, 8], 3, 4, 3],
        [[11, 13, 17, 23, 29, 31, 7, 5, 2, 3], 3, 5, 6],
    ],
)
def test_num_of_subarrays(arr, k, threshold, expected):
    assert Solution().numOfSubarrays(arr, k, threshold) == expected
