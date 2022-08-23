import pytest


from .solution import Solution


@pytest.mark.parametrize("nums, expected", [([3, 5, 2, 3], 7), ([3, 5, 4, 2, 4, 6], 8)])
def test_min_pair_sum(nums, expected):
    assert Solution().minPairSum(nums) == expected
