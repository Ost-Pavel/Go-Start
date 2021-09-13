package main

import "fmt"

func main() {
    var n, sum = 4, 1
    fmt.Scanln(&n)
    for x := n; x > 0; x-- {
        sum *= x
    }
    fmt.Println(sum)
}
