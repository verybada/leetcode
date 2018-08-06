class Solution:
    def lemonadeChange(self, bills):
        """
        :type bills: List[int]
        :rtype: bool
        """
        cashes = {
            5: 0,
            10: 0,
            20: 0,
        }
        for bill in bills:
                            
            cashes[bill] += 1
            
            if bill == 5:
                pass
            elif bill == 10:
                if cashes[5] >= 1:
                    cashes[5] -= 1
                    continue
                return False
            elif bill == 20:
                if cashes[10] >= 1 and cashes[5] >= 1:
                    cashes[10] -= 1
                    cashes[5] -= 1
                    continue
                elif cashes[5] >= 5:
                    cashes[5] -= 3    
                    continue
                return False

        return True 
