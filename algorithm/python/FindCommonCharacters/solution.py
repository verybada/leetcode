class Solution:
    def commonChars(self, words: list[str]) -> list[str]:
        char_maps = list()
        for word in words:
            char_map = dict()
            for char in word:
                count = char_map.get(char, 0)
                char_map[char] = count + 1
            char_maps.append(char_map)

        common_chars = list()
        char_map = char_maps[0]
        for char, count in char_map.items():
            for _ in range(count):
                for other_char_map in char_maps[1:]:
                    if char not in other_char_map:
                        break

                    count = other_char_map.pop(char, 0)
                    if count == 0:
                        break
                    elif count > 1:
                        other_char_map[char] = count - 1
                else:
                    common_chars.append(char)

        return common_chars
