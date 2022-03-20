import pytest

from .number_complement import Solution


@pytest.mark.parametrize('num, expected', [
    [5, 2],
    [1, 0],
    [2147483647, 0],
])
def test_number_complement(num, expected):
    assert Solution().findComplement(num) == expected
