package main

import "fmt"

func main() {
	matriz := []int{2, 3, 1}
	fmt.Println(matriz)
	variante2()
	variante3()
	variante4()

}
func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[1:3]
	fmt.Println(porcion)
}
func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Longitud %d Capacidad %d \n", len(elementos), cap(elementos))
}
func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Longitud %d Capacidad %d \n", len(nums), cap(nums))
}
