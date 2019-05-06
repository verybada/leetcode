class Solution(object):
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        if haystack == needle:
            return 0
        
        if not needle:
            return 0
            
        max_haystack = len(haystack)
        max_needle = len(needle)
        
        for i in range(max_haystack):
            for j in range(max_needle):
                if i+j > max_haystack-1:
                    return -1
            if haystack[i+j] == needle[j]:
                if j == max_needle - 1:
                    return i
            else:
                break
        return -1

        #if needle not in haystack:
        #    return -1
        #return haystack.index(needle)
