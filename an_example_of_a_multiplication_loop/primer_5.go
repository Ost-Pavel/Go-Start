package main

import "fmt"

func main() {
    var in int
    fmt.Scan(&in)
    fmt.Println(fac(in))
}

func fac(n int) int {
    if n == 1 || n == 0 {
        return n
    }
    return n * fac(n - 1)
}
