package main

import "fmt"

func main() {
// здесь вам нужно написать свой код
    var number int
    result := 1
    fmt.Scanln(&number)
    for i := 2; i <= number; i++ {
        result *= i
    }
    fmt.Print(result)
}
