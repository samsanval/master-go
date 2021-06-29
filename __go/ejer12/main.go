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
}
