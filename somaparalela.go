package main

import "fmt"

func somaParalelaAux(start int, end int, step int, result chan int) {
    var sun = 0
    for i := start; i < end; i += step {
        sun += i
    }
    result <- sun
}

func somaParalela(n int, qntThread int) int {
    var c = make(chan int, qntThread)

    for i := 0; i < qntThread; i++ {
        go somaParalelaAux(i, n, qntThread, c)
    }

    var s = 0
    for i := 0; i < qntThread; i++ {
        s = s + <-c
    }
    return s
}

func main() {
    fmt.Println(somaParalela(10, 2))
}