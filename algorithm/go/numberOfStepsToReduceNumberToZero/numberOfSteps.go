package main

func numberOfSteps(num int) int {
    if num == 0 {
        return 0
    }
    count := 1
    for num != 1 {
      if(num % 2 == 0 ) {
        num /= 2
      } else {
          num -= 1
      }
        count++
    }
    return count
}
