package main

import "fmt"

func main() {

	var a int
	var b int

	fmt.Println("숫자 입력 : ")
	fmt.Scanln(&a, &b)

	fmt.Println("res = ", a, b)

}
