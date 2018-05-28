package main

import "fmt"

func main(){
    a, b := 10, 20 //declare variable a & b

    fmt.Println("a is", a, " and b is", b)

    d, c := 40, 50 //b ia already declared but c is new

    fmt.Println("b is ", d, " and c is ", c)

    a, b , e := 30,60,90
 
    fmt.Println("a is", a, "b is ", b, " and e is", e)
}
