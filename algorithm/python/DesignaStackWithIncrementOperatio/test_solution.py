from .solution import CustomStack


def test_custom_stack():
    s = CustomStack(3)
    s.push(1)
    s.push(2)
    assert s.pop() == 2
    s.push(2)
    s.push(3)
    s.push(4)
    s.increment(5, 100)
    s.increment(2, 100)
