package permutation

import (
  "testing"
)

/*
func TestPermutator(t *testing.T) {
  p := NewPermutator(2, 2)
  perm, complete := p.NextPermutation()
  if complete {
    t.Error("Permutation should not be complete")
  }
  if len(perm) != 2 {
    t.Error("Permutation given not of correct length.")
  }

  for _, v := range perm {
    if v >= 2 && v < 0 {
      t.Error("Permutation values should be between 0 and 1")
    }
  }

  perm, complete = p.NextPermutation()
  if complete {
    t.Error("Permutation should still not be complete")
  }

  perm, complete = p.NextPermutation()
  if !complete {
    t.Error("Permutation should be complete")
  }
  if len(perm) != 0 {
    t.Error("Permutation should be of length 0.")
  }
}
*/

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
