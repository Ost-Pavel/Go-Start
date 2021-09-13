package main

import (
	"fmt"
)

func fact(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	var n uint
	fmt.Scan(&n)
	fmt.Println(fact(n))

}
