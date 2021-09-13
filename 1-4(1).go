package main
import "fmt"
func main() {

a,b := 1, 1
for a < 5 {
  b *= a
  a++
}
fmt.Println(b)

}
