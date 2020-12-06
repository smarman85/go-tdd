package iteration

import (
  "testing"
  "fmt"
)

func TestRepeat(t *testing.T) {
  repeated := Repeat("a", 6)
  expected := "aaaaaa"

  if repeated != expected {
    t.Errorf("expected %q but got %q", expected, repeated)
  }
}

func ExampleRepeat() {
  repeated := Repeat("b", 9)
  fmt.Println(repeated)
  // Output: bbbbbbbbb
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i ++ {
    Repeat("a", 5)
  }
}
