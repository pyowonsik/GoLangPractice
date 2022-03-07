package main

import (
	"fmt"
)

func getAge() int {
	return 25
}

func main() {

	switch age := getAge(); age {

	case 25:
		fmt.Println("나의 나이")
	default:
		fmt.Println("몰라 !!")

	}

}
