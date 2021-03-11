package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5+7 = %d\n", Calculo(5, 7))
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Resto 7-5 = %d\n", Calculo(7, 5))
	Operaciones()
	tablaDos := 2
	miTabla := Tabla(tablaDos)
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}
}
func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 14
		return a + b
	}
	fmt.Println(resultado())
}

// Closure
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int { //Cuando llamamos a miTabla ejecutamos sólo la función
		secuencia++
		return numero * secuencia
	}

}
