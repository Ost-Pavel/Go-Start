package main

import (
	"fmt"
	"math/rand"
)

func randNumsGenerator(n int) <-chan int {
	ch := make(chan int, 1)
	var a int 
	
	go func(n int) {
		for i := 0; i < n; i ++ {
			a = rand.Intn(10)
			ch <- a	
		}
		close(ch)
	}(10)
	
	return ch
}

func main() {	
	for num := range randNumsGenerator(10) {
		fmt.Println(num)
	}

}
