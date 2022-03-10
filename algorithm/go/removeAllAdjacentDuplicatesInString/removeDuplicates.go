package main

func removeDuplicates(s string) string {
    stack := make([]rune, 0, len(s))
    for _, char := range s {
        length := len(stack)
        if length == 0 {
            stack = append(stack, char)
            continue
        }
        
        lastChar := stack[length-1]
        if lastChar != char {
            stack = append(stack, char)
            continue
        }
        
        stack = stack[:length-1]
    }
    
    return string(stack)
}
