import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "data, expected",
    [
        ["123", "abc"],
        ["10#11#12", "jkab"],
        ["1326#", "acz"],
    ],
)
def test_freq_alphabets(data, expected):
    assert Solution().freqAlphabets(data) == expected


@pytest.mark.parametrize(
    "data, expected",
    [
        ["123", "abc"],
        ["10#11#12", "jkab"],
        ["1326#", "acz"],
    ],
)
def test_freq_alphabets_stack(data, expected):
    assert Solution().freqAlphabets_stack(data) == expected
