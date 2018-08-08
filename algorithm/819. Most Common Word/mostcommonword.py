class Solution:
    def mostCommonWord(self, paragraph, banned):
        """
        :type paragraph: str
        :type banned: List[str]
        :rtype: str
        """
        banned_dict = dict()
        for b in banned:
            banned_dict[b] = 0
        maps = dict()
        repeated_count = 0
        word = None
        tmp = ''
        
        def _estimate_word(w):
            nonlocal banned_dict
            nonlocal maps
            nonlocal word
            nonlocal repeated_count
            
            if w in banned_dict:
              return

            maps.setdefault(w, 0)
            maps[w] += 1

            if w == word:
                repeated_count += 1
                return                  
                
            if maps[w] <= repeated_count:
                return

            word = w
            repeated_count = maps[w]   
            
        for char in paragraph:
            if char.isspace():
                _estimate_word(tmp)
                tmp = ''

            elif char.isalpha():
                tmp += char.lower()
        
        if tmp:
            _estimate_word(tmp)
        return word
