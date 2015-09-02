class Solution(object):
    def _scan(self, s, length_s, left, right):
        checked = False
        offset = 0
        while True:
            left_ptr = left-offset
            right_ptr = right+offset
            if left_ptr < 0 or right_ptr >= length_s:
                offset -= 1
                break

            if s[left_ptr] != s[right_ptr]:
                offset -= 1
                break

            offset += 1
            checked = True

        if checked is False:
            return None
        return s[left-offset:right+offset+1]

    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        length = len(s)
        if length == 1:
            return s

        left = 0
        right = 0
        if length % 2 == 0:
            right = length >> 1
            left = right - 1
        else:
            middle = length >> 1
            left = middle - 1
            right = middle + 1

        longest_palindrom = None
        longest_len = 0

        palindrome = self._scan(s, length, left, right)
        if palindrome != None:
            current_len = len(palindrome)
            if current_len > longest_len:
                longest_palindrom = palindrome
                longest_len = current_len

        offset = 1
        while True:
            if left - offset < 0 or right+offset >= length:
                break

            palindrome = self._scan(s, length, left-offset, right-offset)
            if palindrome != None:
                current_len = len(palindrome)
                if current_len > longest_len:
                    longest_palindrom = palindrome
                    longest_len = current_len

            palindrome = self._scan(s, length, left+offset, right+offset)
            if palindrome != None:
                current_len = len(palindrome)
                if current_len > longest_len:
                    longest_palindrom = palindrome
                    longest_len = current_len

            offset += 1

        return longest_palindrom


if __name__ == "__main__":
    obj = Solution()
    # s = "abba"
    # s = "abc"
    # s = "a"
    # s = "aaaa"
    s = "abb"
    print obj.longestPalindrome(s)
