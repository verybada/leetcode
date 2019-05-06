class Solution(object):
    def isValid(self, s):
        '''
        :type s: str
        :rtype bool
        '''
        mapping = {
            ')': '(',
            '}': '{',
            ']': '['
        }
        candidates = list()
        for c in s:
            if c in mapping.values():
                # append to stack if they are leftparentheses
                candidates.append(c)
                continue

            # right parentheses but stack is empty, invalid
            if not candidates:
                return False

            # parenthese not match
            last = candidates.pop(-1)
            if last != mapping[c]:
                return False

        # stack should be empty or it is invalid
        return not candidates

if __name__ == '__main__':
    assert Solution().isValid('{}')
    assert Solution().isValid('{}')
    assert Solution().isValid('()')
    assert Solution().isValid('[]')
    assert Solution().isValid('[[]]')
    assert Solution().isValid('([])')
    assert Solution().isValid('(') is False
    assert Solution().isValid('([)]') is False
    assert Solution().isValid('(])') is False
    assert Solution().isValid('(]') is False
    assert Solution().isValid('()()]') is False
