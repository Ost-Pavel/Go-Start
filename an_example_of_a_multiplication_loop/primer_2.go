package main

import "fmt"

func main() {
    var number int
    total := 1
    fmt.Scanln(&number)
    for i:= 1; i <= number; i ++ {
        total *= i;
    }
    fmt.Println(total)
}
