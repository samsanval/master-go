package main

import (
	"fmt"
	"time"
)

type User struct {
	id         int
	name       string
	dateInsert time.Time
	status     bool
}

func (this *User) insertUser(id int, name string, dateInsert time.Time, status bool) {
	this.id = id
	this.name = name
	this.dateInsert = dateInsert
	this.status = status
}

type pepe struct {
	User
}

func main() {
	var pepe = new(pepe)
	pepe.insertUser(1, "Pepe", time.Now(), true)
	fmt.Println(pepe.User)
}
