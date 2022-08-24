import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "s, expected",
    [
        ["", 0],
        ["[]", 0],
        ["[][]", 0],
        ["][][", 1],
        ["]]][[[", 2],
    ],
)
def test_min_swaps(s, expected):
    assert Solution().minSwaps(s) == expected
