class Solution:

    MORSE_CODES = {'a': '.-', 'b': '-...', 'c': '-.-.', 'd': '-..', 'e': '.', 'f': '..-.', 'g': '--.', 'h': '....', 'i': '..', 'j': '.---', 'k': '-.-', 'l': '.-..', 'm': '--', 'n': '-.', 'o': '---', 'p': '.--.', 'q': '--.-', 'r': '.-.', 's': '...', 't': '-', 'u': '..-', 'v': '...-', 'w': '.--', 'x': '-..-', 'y': '-.--', 'z': '--..'}

    def _to_morse(self, word):
        codes = list()
        for c in word:
            codes.append(self.MORSE_CODES[c])
            return ''.join(codes)

    def uniqueMorseRepresentations(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        codes = dict()
        for word in words:
            morse = self._to_morse(word)
            if morse in codes:
                continue

            codes[morse] = morse

        return len(codes)
