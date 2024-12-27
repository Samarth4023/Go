package main

import (
    "fmt"
    "sort"
)

func main() {
    numbers := []int{4, 2, 7, 1, 9}
    sort.Ints(numbers)
    fmt.Println("Sorted Numbers:", numbers)
}
