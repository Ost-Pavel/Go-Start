/*На вход подаются три целых числа. Необходимо сосчитать и вывести их сумму и произведение на разных строках. 
Sample Input: 2 3 4  
Sample Output:
9
24 */


package main

import "fmt"

func main() {
	a, b, c := 2, 3, 4
	fmt.Scanln(&a, &b, &c)
	res_sum := a + b + c
	fmt.Println(res_sum)

	res_x := a * b * c
	fmt.Println(res_x)

}
