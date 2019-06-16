package basicCalculatorII

import (
)

func calculate(s string) int {
    numS := make([]int, 0, len(s))
    num := 0
    prev_sign := '+'
    for _, c := range s {
        if c == ' ' {
            continue
        } else if c >= '0' && c <= '9' {
            num = num*10 + int(c - '0')
        } else {
            if prev_sign == '+' {
                numS = append(numS, num)
            } else if prev_sign == '-' {
                numS = append(numS, -num)
            } else if prev_sign == '*' {
                last_index := len(numS)-1
                prev_num := numS[last_index]
                numS[last_index] = prev_num*num
            } else if prev_sign == '/' {
                last_index := len(numS)-1
                prev_num := numS[last_index]
                numS[last_index] = prev_num/num
            }
            num = 0
            prev_sign = c
        }
    }

    if prev_sign == '+' {
        numS = append(numS, num)
    } else if prev_sign == '-' {
        numS = append(numS, -num)
    } else if prev_sign == '*' {
        last_index := len(numS)-1
        prev_num := numS[last_index]
        numS[last_index] = prev_num*num
    } else if prev_sign == '/' {
        last_index := len(numS)-1
        prev_num := numS[last_index]
        numS[last_index] = prev_num/num
    }

    sum := 0
    for _, num := range numS {
        sum += num
    }
    return sum
}
