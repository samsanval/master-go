package main

import (
	"log"
)

func main() {
	/*var file string = "test.txt"
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}*/
	testPanic()
}

func testPanic() {
	defer func() {
		rec := recover()
		if rec != nil {
			log.Fatalf("Error %v", rec)
		}
	}()
	var a int = 1
	if a == 1 {
		panic("Se encontro 1")
	}
}
