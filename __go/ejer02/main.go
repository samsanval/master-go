package main

import (
	"fmt"
	"strconv"
)

var numero int
var text string
var status bool

func main() {
	numero2, numero3, numero4, texto := 2, 5, 6, "esto es un texto"
	var status bool = true
	var moneda int64 = 45
	text = fmt.Sprintf("%d", moneda)
	text = strconv.Itoa(int(moneda))
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)
	fmt.Println(text)
	showStatus()
}
func showStatus() {
	fmt.Println(status)
}
