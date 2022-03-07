package main

import "fmt"

// 정의
type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Black
)

func colorToString(color ColorType) string {

	switch color {

	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Black:
		return "Black"
	default:
		return ".."
	}

}

func myFavoriteColor() ColorType {

	return Black
}

func main() {

	fmt.Println("my favorite color is ", colorToString(myFavoriteColor()))

}
