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
        >>> obj.lengthOfLongestSubstring("pwwkew")
        3
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
        for char in s:
            if char in cache:
                maxlen = max(maxlen, index-start)
                # make sure next start position not less than current position
                start = cache[char]+1 if cache[char]+1 >= start else start

            cache[char] = index
            index += 1

        return max(maxlen, index-start)
