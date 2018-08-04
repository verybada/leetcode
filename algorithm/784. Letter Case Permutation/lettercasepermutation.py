class Solution:
    def letterCasePermutation(self, S):
        """
        :type S: str
        :rtype: List[str]
        """
        results = list()
        S_len = len(S)
        
        def backtracking(index, path):            
            nonlocal S
            nonlocal S_len
            nonlocal results
            
            if index == S_len:
                results.append(path)
                return
            
            c = S[index]
            
            index += 1
            if c.isalpha():
                backtracking(index, path+c.lower())
                backtracking(index, path+c.upper())
            else:
                backtracking(index, path+c)
                
        backtracking(0, '')
        return results
