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
        # the variable which keeps our start position, it will be reset after we get duplicated character
        start = 0
        # the variable which keeps current index
        index = 0
        # create cache, Key is char, Value is index of this char
        cache = dict()
        for char in s:
            # if this character is occupied, means the longest distance is (current position - start position)
            if char in cache:
                # get longest
                maxlen = max(maxlen, index-start)
                # make sure next start position not less than current position
                start = cache[char]+1 if cache[char]+1 >= start else start

            # put into cache
            cache[char] = index
            index += 1

        return max(maxlen, index-start)
