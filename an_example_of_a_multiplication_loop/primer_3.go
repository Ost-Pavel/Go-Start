package main

import "fmt"

func main() {
    var number int
    fmt.Scanln(&number) 
    factorial := 1
    for number != 0{
        factorial *= number
        number--;
    }
    fmt.Println(factorial) 
    
}
