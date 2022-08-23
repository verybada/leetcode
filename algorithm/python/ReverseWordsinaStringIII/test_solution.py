import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "data, expected",
    [
        ("Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"),
        ("God Ding", "doG gniD"),
        ("A", "A"),
    ],
)
def test_reverse_words(data, expected):
    assert Solution().reverseWords(data) == expected
