class CustomStack:
    def __init__(self, max_size: int):
        self._max_size = max_size
        self._stack = list()

    def is_full(self):
        return len(self._stack) == self._max_size

    def is_empty(self):
        return not self._stack

    def push(self, x: int) -> None:
        if self.is_full():
            return
        self._stack.append(x)

    def pop(self) -> int:
        if self.is_empty():
            return -1
        return self._stack.pop(-1)

    def increment(self, k: int, val: int) -> None:
        last = len(self._stack)
        if k > last:
            k = last

        for i in range(k):
            self._stack[i] += val


# Your CustomStack object will be instantiated and called as such:
# obj = CustomStack(maxSize)
# obj.push(x)
# param_2 = obj.pop()
# obj.increment(k,val)
