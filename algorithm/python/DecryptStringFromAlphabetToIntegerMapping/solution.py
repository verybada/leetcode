class Solution:
    def freqAlphabets(self, s: str) -> str:
        result = list()
        index = len(s) - 1
        while index >= 0:
            char = s[index]
            if char == "#":
                value = int(s[index - 2 : index])
                char = self._int_to_char(value)
                result.append(char)
                index -= 3
            else:
                char = self._int_to_char(int(char))
                result.append(char)
                index -= 1
        result = result[::-1]
        return "".join(result)

    def freqAlphabets_stack(self, s: str) -> str:
        stack = list()
        for char in s:
            if char == "#":
                value = stack.pop() + stack.pop() * 10
                stack.append(value)
                continue

            stack.append(int(char))

        for index, value in enumerate(stack):
            stack[index] = self._int_to_char(value)
        return "".join(stack)

    def _int_to_char(self, value: int) -> str:
        return chr(97 + value - 1)
