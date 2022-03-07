package main

import "fmt"

func main() {

	var a int = 3
	var b float64 = 23.4
	var c int64 = 234

	var res1 = float64(a) + b
	var res2 = int64(b) + c
	var res3 = int64(a) + c

	fmt.Println(res1, res2, res3)

}
