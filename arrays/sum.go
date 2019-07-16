package sum

func sum(numbers []int) (sum int) {
  for _, number := range numbers {
    sum += number
  }
  return
}