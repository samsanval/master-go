package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numero1 int
	var numero2 int
	var resultado int
	var leyenda string
	fmt.Println("Introduzca un numero")
	fmt.Scanf("%d", &numero1)
	fmt.Println("Introduzca el segundo numero")
	fmt.Scanf("%d", &numero2)
	fmt.Println("Introduzca la leyenda")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
