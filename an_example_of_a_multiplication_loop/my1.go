package main
import ("fmt")
func main() {
	
b:=4
a:=1
fmt.Scanln(&b)
for i := 1; i <= b; i++ {
a*=i
  }
fmt.Println(a)
}
