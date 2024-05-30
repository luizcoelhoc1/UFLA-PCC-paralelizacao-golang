package main

import (
    "fmt"
    "time"
)

func x(i int, c chan int) {
    fmt.Println(<-c)
    c <- i
}
func main() {
    var qntThread = 3
    var k = make(chan int, qntThread+1)
    k <- 5
    for i := 0; i < qntThread; i++ {
        go x(i, k)
    }
    time.Sleep(2000 * time.Millisecond)
    fmt.Println(<-k)
}