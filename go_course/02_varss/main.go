package main

import "fmt"

func main() {
	var name string = "Brad"
	var age int = 18
	const iscool = true

	fmt.Println(name, age)
	fmt.Printf("%T\n", iscool)
}
