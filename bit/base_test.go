package bit

import (
	"fmt"
	"testing"
)

func TestEven(t *testing.T) {
	array := Even(10)
	for _, v := range array {
		println(v)
	}
}

func TestSwap(t *testing.T) {
	a, b := 3, 7
	a, b = Swap(a, b)
	fmt.Printf("a=%d b=%d", a, b)
}

func TestShifting(t *testing.T) {
	a := Shifting(10)
	fmt.Printf("%d", a)
}

func TestTransformation(t *testing.T) {
	a := Transformation(10)
	fmt.Printf("%d", a)
}

func TestGetMaxNum(t *testing.T) {
	a := GetMaxNum(-13, 12)
	fmt.Printf("%d", a)
}
