package simplifypath

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	previousSlashIndex := 0
	handleValue := func(value string) {
		switch value {
		case "", ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, value)
		}
	}

	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			value := path[previousSlashIndex+1 : i]
			handleValue(value)

			previousSlashIndex = i
		}
	}

	if previousSlashIndex != len(path)-1 {
		value := path[previousSlashIndex+1:]
		handleValue(value)
	}

	return "/" + strings.Join(stack, "/")
}
