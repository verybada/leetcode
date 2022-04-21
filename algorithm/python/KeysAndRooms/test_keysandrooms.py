import pytest

from .keysandrooms import Solution


@pytest.mark.parametrize('rooms,expected', [
    ([[1],[2],[3],[]], True),
    ([[1,3],[3,0,1],[2],[0]], False),
    ([[3], [], [1], [2]], True),
    ([[], [0], [1], [2]], False),
    ([[1,2,3], [0], [1], [2]], True),
])
def test_keys_and_rooms(rooms, expected):
    result = Solution().canVisitAllRooms(rooms)
    assert result == expected
