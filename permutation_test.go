package permutation

import (
  "testing"
)

func TestPermutatorNumber(t *testing.T) {
  ch := make(chan []int)
  Permute(5, 5, ch)

  i := 0
  for _ = range ch {
    i++
  }
  if i != 120 {
    t.Error("Permutations for Max 5, Length 5 should equal 120.")
  }
}
