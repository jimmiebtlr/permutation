// Package permutation provides concurrent permutation enumeration.
package permutation

// Permute fills a channel with permutations of length, and at most value max.
// It launches a goroutine to concurrently handle the filling of the channel.
// The channel is intended to be used as the index in an array.
func Permute( max int, length int, ch chan []int){
  go fillPermChan( length, max, ch)
}

// fillPermChan fills a channel with permutations of length, and at most value max.
func fillPermChan( length, max int, ch chan []int ){
  used := make( []bool, max )
  perm := initialPerm( length, max, used )

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

// initialPerm takes max and length, creats the initial permutation.
// Used is modified to reflect the state of the initial permutation.
func initialPerm( length, max int, used []bool )( perm []int ){
  perm = make( []int, length )
  for i := int(0); i < length; i++ {
    perm[i] = length - i - 1
    used[length-i-1] = true
  }
  return perm
}

// incrementElement updates perm to the next avaliable.
// It ensures each element is unique using used, and updates it to reflect changes.
// Recursivly calls itself if the current index cannot be incremented further.
// Returns success when the next permutation exists, false if not (last perm was last possible permutation).
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

// nextEmptyElement finds the next unused element from used.
func nextEmptyElement(current, max int, used []bool) (nxtEmpty int, increase bool) {
  for i := current + 1; i < max; i++ {
    if used[i] == false {
      return i, true
    }
  }
  return 0, false
}
