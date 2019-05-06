import re


class Solution:

    def toGoatLatin(self, s):
        """
        :type S: str
        :rtype: str
        """
        prog = re.compile(r'^[aeiou]', re.IGNORECASE)
        index = 1
        new_string = list()
        for word in s.split():
            if prog.match(word):
                word = word + 'ma' + 'a'*index
            else:
                word = word[1:] + word[0] + 'ma' + 'a'*index

            new_string.append(word)
            index += 1

        return ' '.join(new_string)
