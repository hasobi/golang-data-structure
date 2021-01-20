package main

import "fmt"

var p = 10
var g = &p

var q = *p

func main() {
	fmt.Println("p is", p)
	fmt.Println("addres p is ", g)
	// fmt.Println(q)
}
