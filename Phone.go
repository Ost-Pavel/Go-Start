package main

import	"fmt"

func main() {

	
x:=1100
switch {
  case x > 1000:
    fmt.Println("Apple")
  case x>=500 && x<= 1000:
    fmt.Println("Samsung")
  case x <= 500:
    fmt.Println("Nokia с фонариком")

/*  На вход подается целое число, сумма денег, которые у вас есть. Ваша задача - вывести марку телефона, которую вы можете себе позволить купить.

Если сумма больше 1000 - вывести Apple
Если сумма от 500 до 1000 (включительно) - вывести Samsung
Если сумма меньше 500 - вывести Nokia с фонариком */

}
}
