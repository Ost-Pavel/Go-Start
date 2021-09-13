package main

import "fmt"

func main() {
	sum := 1
	 var num int 
    fmt.Scanln(&num)

	for num > 0 {
		sum = sum * num
		num--
	}
	fmt.Println(sum)
}
