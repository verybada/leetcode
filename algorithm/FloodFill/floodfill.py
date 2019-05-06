class Solution:
    def _fill(self, image, sr, sc, new_color, expected_color, max_row, max_column):
        orig_color = image[sr][sc]
        if orig_color != expected_color:
            return 
        
        image[sr][sc] = new_color
        if sr >= 1:
            self._fill(image, sr-1, sc, new_color, expected_color, max_row, max_column)
        if sr+1 < max_row:
            self._fill(image, sr+1, sc, new_color, expected_color, max_row, max_column)
        if sc >= 1:
            self._fill(image, sr, sc-1, new_color, expected_color, max_row, max_column)
        if sc+1 <= max_column:
            self._fill(image, sr, sc+1, new_color, expected_color, max_row, max_column)
        return image
       
    def floodFill(self, image, sr, sc, new_color):
        """
        :type image: List[List[int]]
        :type sr: int
        :type sc: int
        :type newColor: int
        :rtype: List[List[int]]
        """
        old_color = image[sr][sc]
        max_row = len(image)
        max_column = len(image[0])
        self._fill(image, sr, sc, new_color, old_color, max_row, max_column)
        return image
