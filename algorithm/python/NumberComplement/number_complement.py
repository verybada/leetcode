class Solution:
    def findComplement(self, num: int) -> int:
        n = 0
        while n<32:
            value = 2**n
            print(f'{n} {value} {num}')
            if num < value:
                return num ^ (value-1)

            n += 1

        assert 0
