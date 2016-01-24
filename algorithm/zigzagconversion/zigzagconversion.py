class Solution(object):
    def _get_list(self, _lists, index):
        return _lists[index]

    def _update_ptr(self, direction, ptr, _max, _min):
        ptr = ptr + (direction*1)
        if ptr > _max:
            ptr = _max - 1
            direction = -1
        elif ptr < _min:
            ptr = 1
            direction = 1

        return direction, ptr

    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        # it is not necessary doing convert if numRows is 1
        if numRows == 1:
            return s

        _s_lists = list()
        for i in xrange(0, numRows):
            _s_lists.append(list())

        direction = 1
        list_ptr = 0
        for i in xrange(0, len(s)):
            _l = self._get_list(_s_lists, list_ptr)
            _l.append(s[i])

            direction, list_ptr = self._update_ptr(direction, list_ptr, numRows-1, 0)

        new_s = ""
        for i in xrange(0, numRows):
            new_s += "".join(_s_lists[i])

        return new_s

if __name__ == "__main__":
    a = Solution()
    print a.convert("ABC", 1)
    print a.convert("PAYPALISHIRING", 3)
