package main

import (
    "fmt"
)

func main() {
    var x, y uint16
    fmt.Scan(&x, &y)
    if x % 2 == 0 {
        fmt.Println("Even");
        return
    }
    if y % 2 == 1 {
        fmt.Println("Odd");
        return
    }
    fmt.Println("Even");
}