package removeOutermostParentheses

import (
    "testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
    cases := []struct {
        input string
        output string
    } {
        {
            "()()",
            "",
        },
        {
            "(()())(())",
            "()()()",
        },
        {
            "(()())(())(()(()))",
            "()()()()(())",
        },
    }

    for _, test := range cases {
        result := removeOuterParentheses(test.input)
        if result != test.output {
            t.Fatalf("Input: %s, Expected: %s, Result: %s",
                test.input, test.output, result)
        }
    }
}
