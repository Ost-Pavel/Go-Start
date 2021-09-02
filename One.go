


num := 3
 
if num == 1 {
  fmt.Println("Один")
} else if num == 2 {
  fmt.Println("Два")
} else  if num == 3 {
  fmt.Println("Три")
} else {
  fmt.Println(num)
}

// Код проверяет значение переменной num и выводит соответствующий текст.
// Мы можем добиться того же результата с помощью оператора switch:

num := 3

switch num {
  case 1:
    fmt.Println("One")
  case 2:
    fmt.Println("Two")
  case 3:
    fmt.Println("Three")
  default:
    fmt.Println(num)
}

/* Оператор switch делает код короче и легче для чтения.
Каждый оператор case включает в себя значение, с которым нужно сравнить, и код, который нужно выполнить после двоеточия.
Простой способ заменить цепочку if/else - использовать switch без выражения. Таким образом, каждый оператор case может включать условие: */

x := 2

switch {
  case x > 0 && x < 10:
    fmt.Println("something")
  case x > 10:
    fmt.Println("something else")
}


// Далее Пример
package main

import "fmt"

func main() {
    var gender int
    ______ gender {
        case 1:
        fmt.Println("Мужчина")
        ____ 2:
        fmt.Println("Женщина")
        default_
        fmt.Println("Еще не определился")
    }
}
