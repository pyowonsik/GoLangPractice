package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	var x int = 123
	var y int = 456

	fmt.Println("a : ", a, "b : ", b, "f : ", f) // ==>  %v
	fmt.Printf("a : %v b : %v f : %v\n", a, b, f)

	fmt.Printf("x : %05d y : %05d\n", x, y)
}
