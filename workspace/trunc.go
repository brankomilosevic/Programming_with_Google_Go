package main

import "fmt"

func main() {
    var f float64
    var i int32
    
    fmt.Printf("Enter float number: ")

    _, err := fmt.Scan(&f)

    if err != nil {
        fmt.Println("error scanning value:", err)
    } else {
    	i = int32(f)
        fmt.Printf("Truncated input value is: %d", i)
    }
}