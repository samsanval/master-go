package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(1)
	fmt.Println(numero, estado)
	fmt.Println(calculo(1, 4))
	fmt.Println(calculo(4, 10, 5, 6))

}
func uno(numero int) (z int) {
	z = numero * 2
	return
}
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}

//Parametros variables
func calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total += num
		fmt.Printf("Item %d \n", item)
	}
	return total
}
