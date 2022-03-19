import pytest

from .find_the_highest_altitude import Solution


@pytest.mark.parametrize("gain,expected", [
    [[-5,1,5,0,-7], 1],
    [[0], 0],
    [[-5,-4,-3], 0],
])
def test_find_the_highest_altitude(gain, expected):
    result = Solution().largestAltitude(gain)
    assert result == expected
