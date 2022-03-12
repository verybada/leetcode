package main

func minOperations(logs []string) int {
	steps := 0
	for _, next := range logs {
		switch next {
		case "./":
			continue
		case "../":
			if steps == 0 {
				continue
			}
			steps--
		default:
			steps++
		}
	}
	return steps
}
