package permutation

func Permute( max int, length int, ch chan []int){
  go fillPermChan( length, max, ch)
}


func fillPermChan( length, max int, ch chan []int ){
  used := make( []bool, max )
  perm := initialPerm( max, length, used )

  p := make( []int, length )
  copy(p, perm)
  ch <- p

  for incrementElement( 0, length, max, used, perm ){
    p := make( []int, length )
    copy(p, perm)
    ch <- p
  }
  close(ch )
}

func initialPerm( max, length int, used []bool )( perm []int ){
  perm = make( []int, length )
  for i := int(0); i < length; i++ {
    perm[i] = length - i - 1
    used[length-i-1] = true
  }
  return perm
}

func incrementElement(index, length, max int, used []bool, perm []int) (success bool) {
  // If the value is as high as it can go
  next, inc := nextEmptyElement(perm[index], max, used )
  if inc {
    used[next] = true
    used[perm[index]] = false
    perm[index] = next
  } else {
    // if the final element
    if index == length-1 {
      if max-1 == perm[index] {
        return false
      } else {
        used[perm[index]] = false
        perm[index]++
      }
    } else { // if not the final element
      used[perm[index]] = false
      if !incrementElement(index + 1, length, max, used, perm) {
        return false
      } else {
        if used[0] {
          next, _ = nextEmptyElement(0, max, used)
          perm[index] = next
          used[next] = true
        } else {
          perm[index] = 0
          used[0] = true
        }
      }
    }
  }
  return true
}

func nextEmptyElement(current, max int, used []bool) (nxtEmpty int, increase bool) {
  for i := current + 1; i < max; i++ {
    if used[i] == false {
      return i, true
    }
  }
  return 0, false
}
