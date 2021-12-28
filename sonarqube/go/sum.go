package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sum(2, 2))
}

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mult(a int, b int) int {
	return a * b
}

func Div(a int, b int) int {
	return a / b
}

func Power(a float64, b float64) float64 {
	return math.Pow(a, b)
}

func SayHello() string {
	return "Hello"
}
