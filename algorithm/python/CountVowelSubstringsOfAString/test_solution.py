import pytest

from .solution import Solution


@pytest.mark.parametrize('word, expected', [
    # ['aeio', 0],
    # ['aeiouu', 2],
    # ['unicornarihan', 0],
    # ['cuaieuouac', 7],
    # ['ughspuuoaaaoieiuiaoiuee', 76],
    # ['vmvmxbmcnhqejkplswrmzcikhaiaiiueiuoiuauoaiaoiuoaooljzaculyiyqeshweyqbbdtfzkueuiiuooeeeauooaiiaaaiuou', 224],
    ['iooiuaioaaeoaoiauiuoeioeoueuaeuoeeeeeaiauiiioeouaiouaaiaeeuoeeoeeuuaaeuiueaeeeeiiooeuaoiuoeiooiuaoiu', 4166],
    ['aaaeiouaa', 11],
])
def test_solution(word, expected):
    assert Solution().countVowelSubstrings(word) == expected
