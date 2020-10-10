package thousandSeperator

import (
	"fmt"
	"strings"
)

func thousandSeparator(n int) string {
	value := fmt.Sprintf("%d", n)
	length := len(value)
	output := make([]string, length+(length-1)/3)
	outputIndex := len(output) - 1
	count := 3
	for i := length - 1; i >= 0; i-- {
		output[outputIndex] = string(value[i])
		outputIndex -= 1
		count -= 1
		if outputIndex < 0 {
			break
		}
		if count == 0 {
			output[outputIndex] = "."
			outputIndex -= 1
			count = 3
		}
	}

	return strings.Join(output, "")
}
