package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Scanf("%d", &n) 

    results := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &results[i])
    }

    max := math.MinInt
    for _, score := range results {
        if score > max {
            max = score 
        }
    }

    fmt.Printf("%d\\n", max) 
}