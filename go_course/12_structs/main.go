package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) greet() string {
	return "Hello my name is " + p.name + " " + "and I am " + strconv.Itoa(p.age)
}

func (p *Person) hasbirthday() {
	p.age++
}

func main() {
	person1 := Person{name: "Samatha", age: 22, gender: "female"}
	// person1 := Person{"Samatha",22,"female"}

	person1.hasbirthday()
	fmt.Println(person1.greet())
}
