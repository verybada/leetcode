class Solution:
    def isOneBitCharacter(self, bits):
        """
        :type bits: List[int]
        :rtype: bool
        """
        if bits == [0]:
          return True

        bitlen = len(bits)
        index = 0
        while True:
            offset = 0
            if bits[index] == 0:
                offset = 1
            else:
                offset = 2
                
            if index + offset >= bitlen:
                return False
            elif index + offset == bitlen - 1 and bits[index+offset] == 0:
                return True

            index += offset
