import math

class Solution:
    def addDigits(self, num):
        """
        :type num: int
        :rtype: int
        """
        
        
        # Advance solution
        if num == 0:
            return 0
        elif num % 9 == 0:
            return 9
        else:
            return num % 9
        
        
        def sum(n):
            if n == 0:
                return 0
            elif math.log10(n) < 1:
                return n

            s = 0
            for c in str(n):
                s += int(c)
                
            return sum(s)
        return sum(num)
