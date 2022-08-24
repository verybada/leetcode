import pytest

from .solution import Solution


@pytest.mark.parametrize(
    "pushed, popped, expected",
    [
        [[1, 2, 3, 4, 5], [4, 5, 3, 2, 1], True],
        [[1, 2, 3, 4, 5], [4, 3, 5, 1, 2], False],
    ],
)
def test_validate_stack_seq(pushed, popped, expected):
    assert Solution().validateStackSequences(pushed, popped) == expected
