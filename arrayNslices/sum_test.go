package main

import (
  "testing"
  "reflect"
)

func TestSum(t *testing.T) {

  assertCorrectSum := func(t *testing.T, got, want int, numbers []int) {
    t.Helper()
    if got != want {
      t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
  }

  t.Run("collection of 5 numbers", func(t *testing.T) {
    numbers := []int{1,2,3,4,5}
    got := Sum(numbers)
    want := 15
    assertCorrectSum(t, got, want, numbers)
  })

  t.Run("collection of any size", func(t *testing.T) {
    numbers := []int{1,2,3}
    got := Sum(numbers)
    want := 6
    assertCorrectSum(t, got, want, numbers)
  })
}

func TestSumAll(t *testing.T) {
  got := SumAll([]int{1,2}, []int{0,9})
  want := []int{3,9}

  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestAllTails(t *testing.T) {
  got := SumAllTails([]int{1,2}, []int{0,9})
  want := []int{2,9}

  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
  }
}
