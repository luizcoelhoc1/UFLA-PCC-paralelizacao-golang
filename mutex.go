package main


import (
    "fmt"
    "sync"
)

var result int = 1
var mu sync.Mutex

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go Multiply(&wg, 10)
    go Multiply(&wg, 4)
    wg.Wait()
}

// Multiply -
func Multiply(wg *sync.WaitGroup, factor int) {
    defer wg.Done()

    mu.Lock()
    result = result * factor
    fmt.Println("= ", result)
    defer mu.Unlock()
}