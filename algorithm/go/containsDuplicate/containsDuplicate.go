package containsDuplicate

func containsDuplicate(nums []int) bool {
    duplicated := make(map[int]int)
    for _, value := range nums {
        _, ok := duplicated[value]
        if ok {
            return true
        }
        duplicated[value] = 1
    }
    return false
}
