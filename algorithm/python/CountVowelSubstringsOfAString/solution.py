class Solution:
    def countVowelSubstrings(self, word: str) -> int:
        if len(word) < 5:
            return 0

        head = None
        result = 0
        vowels = 'aeiou'
        vowel_counts = dict()
        for i, char in enumerate(word):
            if char not in vowels:
                vowel_counts.clear()
                head = None
                continue

            if head is None:
                head = i

            count = vowel_counts.get(char, 0)
            vowel_counts[char] = count + 1
            while len(vowel_counts) == 5:
                for remaining_char in word[i:]:
                    if remaining_char not in vowels:
                        break

                    result += 1

                char = word[head]
                count = vowel_counts.pop(char)
                count -= 1
                if count > 0:
                    vowel_counts[char] = count
                head += 1

        return result
