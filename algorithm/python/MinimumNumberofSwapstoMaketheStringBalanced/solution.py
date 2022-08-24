class Solution:
    def minSwaps(self, s: str) -> int:
        unclosed_bracket = 0
        for c in s:
            if c == "[":
                unclosed_bracket += 1
            elif c == "]" and unclosed_bracket:
                unclosed_bracket -= 1

        return (unclosed_bracket + 1) // 2
