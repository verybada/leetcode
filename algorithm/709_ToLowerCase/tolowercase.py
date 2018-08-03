class Solution:
    def toLowerCase(self, str):
        """
        :type str: str
        :rtype: str
        """
        r = list()
        for c in str:
            code = ord(c)
            if code >= 65 and code <= 90:
                r.append(chr(code+32))
            else:
                r.append(c)


        return ''.join(r)
        #return str.lower()

