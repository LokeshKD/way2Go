package main
import "fmt"

func f1(a,b,c int)int{  // taking three parameters and returning their sum
  return a+b+c
}

func f2(a,b int)(int, int , int){ // taking two parameters and returning their sum, difference and product
  n1 := a+b
  n2 := a-b
  n3 := a*b
  return n1,n2,n3
}

func main(){
  fmt.Print(f1(f2(20,10)))  // function call within a function call
}
