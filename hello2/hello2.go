package main

import (
	"bufio"
	"fmt"
	"os"
)

func mul(a int, b int) (int, bool) {
	if a == 0 || b == 0 {
		return 0, false
	} else {
		return a * b, true
	}
}

func nearOne(n int) {

	if n == 0 {
		return
	} else {

		fmt.Print(n, " ")
		n = n - 1
		nearOne(n)
	}

}

func main() {

	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	fmt.Print(" 곱셈할 숫자를 입력 : ")

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {

		c, res := mul(a, b)
		fmt.Println(n, "곱셈 결과 : ", c, " res : ", res)
	}

	fmt.Print("숫자 입력 : ")
	var number int
	fmt.Scanln(&number)

	nearOne(number)

}
