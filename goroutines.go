package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 0; i < 5; i++ {
        fmt.Println(msg)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go printMessage("Hello from Goroutine!")
    printMessage("Hello from Main!")
}
