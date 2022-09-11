class Solution:
    def sortSentence(self, s: str) -> str:
        parts = s.split()
        result = [None] * len(parts)
        for part in parts:
            index = int(part[-1]) - 1
            sentence = part[:-1]
            result[index] = sentence
        return " ".join(result)
