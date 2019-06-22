package removeOutermostParentheses

import (
    "strings"
)


func removeOuterParentheses(S string) string {
    stack := make([]string, 0, len(S))

    ref := 0
    for _, c := range S {
        if c == '(' {
            if ref > 0 {
                stack = append(stack, string(c))
            }
            ref += 1
        } else if c == ')' {
            if ref > 1 {
                stack = append(stack, string(c))
            }
            ref -= 1
        }
    }
    return strings.Join(stack, "")
}
