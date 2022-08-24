import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "words, expected",
    [
        [["bella", "label", "roller"], ["e", "l", "l"]],
        [["cool", "lock", "cook"], ["c", "o"]],
        [["a", "b", "c"], []],
    ],
)
def test_common_chars(words, expected):
    assert Solution().commonChars(words) == expected
