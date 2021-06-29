package main

import "fmt"

type lifeBeing interface {
	isAlive() bool
}

type human interface {
	breath()
	think()
	eat()
	gender() string
}

type animal interface {
	breath()
	eat()
	isCarnivore() bool
}

type vegetal interface {
	vegetalClassification() string
}

/*Human*/
type man struct {
	age         int
	height      float32
	weight      float32
	isBreathing bool
	isThinking  bool
	isEating    bool
	isMan       bool
	alive       bool
}
type woman struct {
	man
}

func (m *man) breath() { m.isBreathing = true }
func (m *man) think()  { m.isThinking = true }
func (m *man) eat()    { m.isEating = true }
func (m *man) gender() string {
	if m.isMan == true {
		return "Man"
	} else {
		return "Woman"
	}
}

func (m *man) isAlive() bool {
	return m.alive
}

func (w *woman) gender() string { return "Woman" }

func isHumanBreathing(hu human) {
	hu.breath()
	fmt.Printf("I am a %s and I am breathing \n", hu.gender())
}

/*Animals*/

type dog struct {
	isBreathing bool
	isEating    bool
	carnivore   bool
	alive       bool
}

func (d *dog) breath()           { d.isBreathing = true }
func (d *dog) eat()              { d.isEating = true }
func (d *dog) isCarnivore() bool { return d.carnivore }

func (d *dog) isAlive() bool {
	return d.alive
}

func isAnimalsBreathing(a animal) {
	a.breath()
	fmt.Printf("I am animal and I am breathing \n")
}
func isAnimalsCarnivore(a animal) int {
	if a.isCarnivore() == true {
		return 1
	}
	return 0
}

func isAlive(v lifeBeing) bool {
	return v.isAlive()
}

func main() {
	var peter = new(man)
	peter.isMan = true
	peter.alive = true
	isHumanBreathing(peter)
	var maria = new(woman)
	isHumanBreathing(maria)

	totalCarnivores := 0
	var dogo = new(dog)
	dogo.carnivore = true
	dogo.alive = true
	isAnimalsBreathing(dogo)
	totalCarnivores += isAnimalsCarnivore(dogo)
	fmt.Println(totalCarnivores)
	fmt.Printf("Los humanos estan vivos %t y los perros estan vivos %t", isAlive(peter), isAlive(dogo))
}
