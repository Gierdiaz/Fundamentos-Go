package main
import ("fmt")

var x bool

func main() {
  fmt.Println(x) //false value

  x = true
  fmt.Println(x) //true value

  x := (10 == 100)
  y := (10 < 100)
  z := (10 > 100)
  
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
}