package iteration

func Repeat(ch string, count int) string {
  var str string
  for i := 0; i < count; i++ {
    str += ch
  }
  return str
}