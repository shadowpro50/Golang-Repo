package main

import (
	"fmt"
	"time"
)

func main() {
	tf := time.Now().Format("02-01-2006")
	fmt.Println(tf)
}
