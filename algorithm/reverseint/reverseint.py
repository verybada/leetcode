class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        >>> s = Solution()
        >>> s.reverse(0)
        0
        >>> s.reverse(123)
        321
        >>> s.reverse(-123)
        -321
        >>> s.reverse(1534236469)
        0
        """
        # int = -2147483648 - 2147483647
        max_value = 2147483647
        neg = False if x >= 0 else True
        if neg:
            max_value += 1
            x = -x

        # convert value
        reverse_x = int(str(x)[::-1])
        # return 0 if overflow
        if reverse_x > max_value:
            return 0

        if neg:
            reverse_x = -reverse_x
        return reverse_x
