package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com\n"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)
}
