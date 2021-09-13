package main

import "fmt"

func main() {
	var a, i, numbers int64
	fmt.Scan(&a)
	numbers = 1
	for i = a; i >= 1; i-- {
		numbers *= i
	}
	fmt.Println(numbers)
}
