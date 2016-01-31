class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        # it is not necessary doing convert if numRows is 1 or numRows > string
        if numRows == 1 or numRows > len(s):
            return s

        _s_lists = [''] * numRows

        direction = 1
        list_ptr = 0
        for i in xrange(0, len(s)):
            _s_lists[list_ptr] += s[i]
            if list_ptr == 0:
                direction = 1
            elif list_ptr == numRows-1:
                direction = -1

            list_ptr += direction

        return ''.join(_s_lists)

if __name__ == "__main__":
    a = Solution()
    print a.convert("ABC", 1)
    print a.convert("ABC", 2)
    print a.convert("PAYPALISHIRING", 3)
