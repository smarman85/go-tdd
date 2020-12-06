package iteration

const repeatCount = 5

func Repeat(str string, repeat int) string {
  var repeated string
  for i := 0; i < repeat; i ++ {
    repeated += str
  }
  return repeated
}
