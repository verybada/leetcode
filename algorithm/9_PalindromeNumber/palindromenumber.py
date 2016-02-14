class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False

        if x < 10:
            return True

        reverse_x = int(str(x)[::-1])
        if reverse_x > 2147483647:
            return False

        if x == reverse_x:
            return True
        return False
