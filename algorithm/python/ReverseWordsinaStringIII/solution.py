class Solution:
    def reverseWords(self, s: str) -> str:
        parts = s.split()
        for index, part in enumerate(parts):
            parts[index] = part[::-1]
        return " ".join(parts)
