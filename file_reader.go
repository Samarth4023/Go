package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("example.txt")
    if err != nil {
        fmt.Println("File reading error:", err)
        return
    }
    fmt.Println("File content:", string(data))
}
