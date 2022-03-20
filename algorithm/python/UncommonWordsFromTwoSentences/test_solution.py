import pytest

from .solution import Solution


@pytest.mark.parametrize('s1, s2, expected', [
    ['this apple is sweet', 'this apple is sour', ["sweet","sour"]],
    ['apple apple', 'banana', ['banana']],
    ['aa', 'bb', ['aa', 'bb']],
    ['a a a',  'a bb', ['bb']],
])
def test_solution(s1, s2, expected):
    assert Solution().uncommonFromSentences(s1, s2) == expected
