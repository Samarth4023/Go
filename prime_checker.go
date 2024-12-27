package main

import "fmt"

func isPrime(num int) bool {
    if num < 2 {
        return false
    }
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    num := 29
    fmt.Printf("%d is prime: %v\n", num, isPrime(num))
}
