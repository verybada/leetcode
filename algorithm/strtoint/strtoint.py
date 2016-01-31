import re


class Solution(object):
    def myAtoi(self, _str):
        """
        Consider str is a garbage, positive, negative or overflow
        :type str: str
        :rtype: int
        >>> s = Solution()
        >>> s.myAtoi("")
        0
        >>> s.myAtoi("                ")
        0
        >>> s.myAtoi("100")
        100
        >>> s.myAtoi("      -111")
        -111
        >>> s.myAtoi("-111")
        -111
        >>> s.myAtoi("+")
        0
        >>> s.myAtoi("-")
        0
        >>> s.myAtoi("+500")
        500
        >>> s.myAtoi("+10abc10")
        10
        >>> s.myAtoi("+aaabbbcccc10")
        0
        >>> s.myAtoi("+aaabbbcccc10a10")
        0
        >>> s.myAtoi("+-10")
        0
        """
        _str = _str.strip()
        _value = ""
        neg = False
        sign = False
        max_value = 2147483647
        for c in _str:
            if c == '-':
                if sign:
                    return 0
                sign = True
                neg = True
                max_value += 1
            elif c == '+':
                if sign:
                    return 0
                sign = True
            elif c.isdigit():
                _value += c
            else:
                break

        if not _value:
            return 0

        _value = int(_value)
        if _value > max_value:
            _value = max_value

        if neg:
            return -_value
        else:
            return _value
