package main

import "fmt"

func main() {
    var numb int
    fmt.Scanln(&numb)
    a := 1
    b := 1
    for b <= numb {
        a *= b
        b++
    }
    fmt.Println(a)
}
