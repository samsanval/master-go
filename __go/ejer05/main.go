package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Printf("Valor de i: %d \n", i)
		if i == 5 {
			i = i * 2
			continue
		}
		i++
	}
	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Introduzca un numero")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var j int = 0
RUTINA:
	for j < 10 {
		if j == 4 {
			j = j + 2
			fmt.Println("Voy a RUTINA")
			goto RUTINA
		}
		fmt.Printf("Valor de j: %d \n", j)
		j++
	}
}
