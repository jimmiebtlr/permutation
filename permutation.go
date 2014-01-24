package permutation

// Max is probably not the best term
// Basically its the largest value minus one for any element
type Permutator struct {
  perm   []uint8
  max    uint8
  length uint8
  used   []bool
  first  bool
}

func NewPermutator(max uint8, length uint8) (p Permutator) {
  p.max = max
  p.perm = make([]uint8, length)
  p.length = length
  // tracks whether the index value is used or not
  p.used = make([]bool, max)
  for i := uint8(0); i < length; i++ {
    p.perm[i] = length - i - 1
    p.used[length-i-1] = true
  }
  p.first = true
  return p
}

func (p *Permutator) NextPermutation() (perm []uint8, complete bool) {
  if p.first {
    p.first = false
    return p.perm, false
  } else if !p.incrementElement(0) {
    return p.perm, false
  } else {
    return nil, true
  }
}

func (p *Permutator) incrementElement(index uint8) (complete bool) {
  // If the value is as high as it can go
  next, inc := p.nextEmptyElement(p.perm[index])
  if inc {
    p.used[next] = true
    p.used[p.perm[index]] = false
    p.perm[index] = next
  } else {
    // if the final element
    if index == p.length-1 {
      if p.max-1 == p.perm[index] {
        return true
      } else {
        p.used[p.perm[index]] = false
        p.perm[index]++
      }
    } else { // if not the final element
      p.used[p.perm[index]] = false
      if p.incrementElement(index + 1) {
        return true
      } else {
        if p.used[0] {
          next, _ = p.nextEmptyElement(0)
          p.perm[index] = next
          p.used[next] = true
        } else {
          p.perm[index] = 0
          p.used[0] = true
        }
      }
    }
  }
  return false
}

func (p *Permutator) nextEmptyElement(current uint8) (empty uint8, increase bool) {
  for i := current + 1; i < p.max; i++ {
    if p.used[i] == false {
      return i, true
    }
  }
  return 0, false
}
