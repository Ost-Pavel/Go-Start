package main

import "fmt"

func main() {
    var input int
    fmt.Scanln(&input)
    
    for i := input-1; i > 0; i-- {        
        input *= i
    }  
    
    fmt.Println(input)
}
