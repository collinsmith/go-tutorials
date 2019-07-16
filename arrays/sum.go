package sum

func sum(numbers []int) (sum int) {
  for _, number := range numbers {
    sum += number
  }
  return
}

func sumAll(numbers ...[]int) []int {
  len  := len(numbers)
  sums := make([]int, len)

  for i, numbers := range numbers {
    sums[i] = sum(numbers)
  }
  return sums
}

func sumAllTails(numbersToSum ...[]int) []int {
  var sums []int
  for _, numbers := range numbersToSum {
    tail := numbers[1:]
    sums = append(sums, sum(tail))
  }

  return sums
}