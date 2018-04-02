package main 

import "fmt"

//Add ..
func Add(a, b int) int{
  return a+b
}


func main () {
  fmt.Println("Hello Hardproof - " + string(Add(5,5)))
}
