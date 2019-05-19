package groupanagrams

import (
    "strings"
    "sort"
)

func groupAnagrams(strs []string) [][]string {
    hash := map[string][]string{}
    for _, str := range strs {
        parts := strings.Split(str, "")
        sort.Strings(parts)
        
        key := strings.Join(parts, "")
        result, ok := hash[key]
        if !ok {
            result = []string{}
        }
        
        hash[key] = append(result, str)
    }
    
    
    resultLen := len(hash)
    results := make([][]string, 0, resultLen)
    for _, value := range hash {
        results = append(results, value)        
    }
    return results
}
