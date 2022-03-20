from typing import List


class Solution:
    def uncommonFromSentences(self, s1: str, s2: str) -> List[str]:
        words = dict()
        for word in s1.split():
            count = words.get(word, 0)
            words[word] = count +1

        for word in s2.split():
            count = words.get(word, 0)
            words[word] = count +1

        results = list()
        for key, value in words.items():
            if value != 1:
                continue
            results.append(key)
        return results
