```go
func main() {
    var i int
    for i = 0; i < 10; i++ {
        go func() {
            fmt.Println(i)
        }()
    }
    time.Sleep(time.Second)
}
```