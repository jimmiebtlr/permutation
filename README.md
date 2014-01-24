```go
ch := make(chan []int)
Permute(5, 5, ch)

i := 0
for p = range ch {
  fmt.Println( p )
}
```

[TODO]
Add GoDoc compatible comments.
Add README
Improve test coverage.
