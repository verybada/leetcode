import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "data, expected",
    [
        ["is2 sentence4 This1 a3", "This is a sentence"],
        ["Myself2 Me1 I4 and3", "Me Myself and I"],
    ],
)
def test_sort_sentence(data, expected):
    assert Solution().sortSentence(data) == expected
