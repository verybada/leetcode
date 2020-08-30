package generateParenthesis

func generateParenthesis(n int) []string {
	outputChan := make(chan string, 1)
	go func() {
		gen(n, n, 1, "(", outputChan)
		close(outputChan)
	}()

	output := make([]string, 0)
	for {
		parenthese, more := <-outputChan
		if !more {
			break
		}

		output = append(output, parenthese)
	}

	return output
}

func gen(
	maxParenthese int, remainParenthese int, openingParenthese int,
	s string, outputChan chan string,
) {
	if remainParenthese == 0 && openingParenthese == 0 {
		outputChan <- s
		return
	}

	if remainParenthese > 0 && openingParenthese <= maxParenthese {
		gen(
			maxParenthese, remainParenthese, openingParenthese+1, s+"(",
			outputChan)
	}
	if openingParenthese > 0 {
		gen(
			maxParenthese, remainParenthese-1, openingParenthese-1, s+")",
			outputChan)
	}
}
