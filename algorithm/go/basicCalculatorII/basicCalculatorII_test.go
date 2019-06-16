package basicCalculatorII

import (
    "testing"
)

func TestCalculate(t *testing.T) {
    cases := []struct {
        input string
        output int
    } {
        {
            "3+2*2",
            7,
        },
        {
            " 3/2 ",
            1,
        },
        {
            " 3+5 / 2 ",
            5,
        },
        {
            "42",
            42,
        },
        {
            "123+456",
            579,
        },
        {
            "100/1/2",
            50,
        },
    }

    for _, test := range cases {
        t.Logf("Testing %s", test.input)
        output := calculate(test.input)
        if  output != test.output {
            t.Fatalf("%s == %d?", test.input, output)
        }
    }
}
