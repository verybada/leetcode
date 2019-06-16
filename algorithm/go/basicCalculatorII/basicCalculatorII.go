package basicCalculatorII

import (
    "strings"
    "strconv"
)

type stack struct{
    s []int
}

func (s *stack) sum() int {
    num := 0
    for _, value := range s.s {
        num += value
    }
    return num
}

func (s *stack) push(value int) {
    s.s = append(s.s, value)
}

func (s *stack) pop() int {
    l := len(s.s)
    v := s.s[l-1]
    s.s = s.s[:l-1]
    return v
}

func isSymbo(s string) bool {
    switch(s) {
    case "+":
        fallthrough
    case "-":
        fallthrough
    case "*":
        fallthrough
    case "/":
        return true
    default:
        return false
    }
}

func calculate(s string) int {
    // Trim all space
    s = strings.Replace(s, " ", "", -1)

    numS := stack{
        make([]int, 0, len(s)),
    }
    getNum := func(num int, symbo string) int {
        switch(symbo) {
        case "+":
            return num
        case "-":
            return -num
        case "*":
            prev_num := numS.pop()
            return prev_num*num
        case "/":
            prev_num := numS.pop()
            return prev_num/num
        }

        return -1
    }

    chars := strings.Split(s, "")
    prev_i := 0
    prev_symbo := "+"
    for i := 0; i<len(chars); i++ {
        c := chars[i]

        switch(c) {
        case "+", "-", "*", "/":
            num, err := strconv.Atoi(s[prev_i: i])
            if err != nil {
                panic(err)
            }

            num = getNum(num, prev_symbo)
            numS.push(num)
            prev_symbo = c
            prev_i = i+1
            break
        }
    }
    num, err := strconv.Atoi(s[prev_i:])
    if err != nil {
        panic(err)
    }
    num = getNum(num, prev_symbo)
    numS.push(num)
    return numS.sum()
}
