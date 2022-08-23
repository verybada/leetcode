import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "data, expected",
    [
        ("aacaba", 2),
        ("abcd", 1),
        ("a", 0),
        ("aaa", 2),
    ],
)
def test_num_splits(data, expected):
    assert Solution().numSplits(data) == expected
