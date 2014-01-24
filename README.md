```go
ch := make(chan []int)
Permute(5, 5, ch)

i := 0
for p = range ch {
  fmt.Println( p )
}
```

```go
ch := make(chan []int)

Permute(len(someArray), len(someArray), ch)

i := 0
for p = range ch {
  total := 0
  for index := p {
    total += someCalc(someArray[index])
  }
  fmt.Printf( "Total for this permutation was %f", total)
}
```

[TODO]
Add GoDoc compatible comments.
Add README
Improve test coverage.
