package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    user := User{Name: "John", Age: 30, Email: "john@example.com"}
    jsonData, _ := json.Marshal(user)
    fmt.Println("JSON Data:", string(jsonData))

    var decodedUser User
    json.Unmarshal(jsonData, &decodedUser)
    fmt.Println("Decoded User:", decodedUser)
}
