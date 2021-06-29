package main

import "fmt"

func main() {
	exponential(1000)

}

func exponential(number int) {
	if number > 10000000 {
		return
	}
	fmt.Println(number)
	exponential(number * 2)
}
