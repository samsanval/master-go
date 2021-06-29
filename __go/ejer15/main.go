package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go myNameSlow("Samuel Sanchez")
	fmt.Println("Estoy aqui")
	var cadena string
	fmt.Scanln(&cadena)

}

func myNameSlow(name string) {
	characters := strings.Split(name, "")
	for _, char := range characters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(char)
	}
}
