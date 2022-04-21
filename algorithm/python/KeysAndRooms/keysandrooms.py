from typing import List


class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        locked_rooms = dict()
        for index, room in enumerate(rooms):
            if index == 0:
                continue
            locked_rooms[index] = 1

        def _discover_room(index):
            keys = rooms[index]
            for key in keys:
                if not locked_rooms:
                    return

                if key not in locked_rooms:
                    continue

                del locked_rooms[key]
                _discover_room(key)

        _discover_room(0)
        return not locked_rooms
