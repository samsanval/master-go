package main

import "fmt"

func main() {
	var campeonato = map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	campeonato["River Plate"] = 25
	fmt.Println(campeonato)
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)
	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje de : %d \n", equipo, puntaje)
	}
	var scoring, exists = campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe %t \n", scoring, exists)
}
