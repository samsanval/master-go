package main

import "fmt"

var result int

func main() {
	fmt.Println("Start")
	result = operacionesMid(add)(2, 5)
	fmt.Println(result)
	result = operacionesMid(substract)(2, 5)
	fmt.Println(result)
	result = operacionesMid(multiply)(2, 5)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}
func multiply(a, b int) int {
	return a * b
}
func operacionesMid(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Start")
		return f(a, b)
	}
}
