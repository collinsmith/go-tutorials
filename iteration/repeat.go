package iteration

func Repeat(ch string) string {
  var str string
  for i := 0; i < 5; i++ {
    str += ch
  }
  return str
}