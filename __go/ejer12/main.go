package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFile()
	readFile2()
	writeFile()
	writeFile2()
}

func readFile() {
	var file, err = ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(file))
	}
}

func readFile2() {
	var file, err = os.Open("./file.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		var scanner = bufio.NewScanner(file)
		for scanner.Scan() {
			var register string = scanner.Text()
			fmt.Printf("Linea > %s \n", register)
		}
	}
	file.Close()
}
func writeFile() {
	var file, err = os.Create("./newFile.txt")
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Fprintln(file, "Esto es un texto metido en el curso de GO")
	file.Close()
}

func writeFile2() {
	var fileName string = "./newFile.txt"
	if append(fileName, "\n Esta es una segunda linea") == false {
		fmt.Println("Error")

	}
}
func append(filenName string, text string) bool {
	var file, err = os.OpenFile(filenName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error")
		return false
	}
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error al escribir")
		return false
	}
	return true
}
