class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        >>> obj = Solution()
        >>> obj.lengthOfLongestSubstring("abc")
        3
        >>> obj.lengthOfLongestSubstring("abcabd")
        4
        >>> obj.lengthOfLongestSubstring("kwweab")
        4
        >>> obj.lengthOfLongestSubstring("")
        0
        >>> obj.lengthOfLongestSubstring("aa")
        1
        >>> obj.lengthOfLongestSubstring("aaa")
        1
        >>> obj.lengthOfLongestSubstring("aaaa")
        1
        >>> obj.lengthOfLongestSubstring("bdb")
        2
        >>> obj.lengthOfLongestSubstring("c")
        1
        >>> obj.lengthOfLongestSubstring("au")
        2
        >>> obj.lengthOfLongestSubstring("abba")
        2
        >>> obj.lengthOfLongestSubstring("tmmzuxt")
        5
        """
        maxlen = 0
        start = 0
        index = 0
        cache = dict()
        #print "c i s m"
        for char in s:
            if char in cache:
                maxlen = max(maxlen, index-start)
                start = cache[char]+1
            cache[char] = index

            #print char, index, start, maxlen
            index += 1

        maxlen = max(maxlen, index-start)
        return maxlen

if __name__ == "__main__":
    obj = Solution()
    print obj.lengthOfLongestSubstring("abc")
    print obj.lengthOfLongestSubstring("abcabd")
    print obj.lengthOfLongestSubstring("kwweab")
    print obj.lengthOfLongestSubstring("")
    print obj.lengthOfLongestSubstring("aa")
    print obj.lengthOfLongestSubstring("aaa")
    print obj.lengthOfLongestSubstring("bdb")
    print obj.lengthOfLongestSubstring("c")
    print obj.lengthOfLongestSubstring("au")
    print obj.lengthOfLongestSubstring("abba")
    print obj.lengthOfLongestSubstring("tmmzuxt")

