class Solution:
    def validateStackSequences(self, pushed: list[int], popped: list[int]) -> bool:
        stack = list()
        for element in pushed:
            stack.append(element)

            while stack and stack[-1] == popped[0]:
                popped.pop(0)
                stack.pop()

        return not stack
