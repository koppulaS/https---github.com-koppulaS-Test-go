package main

import "fmt"

func add(x, y int) int {
	for i := 0; i < y; i++ {
		x++
	}
	return x
}

var a, b int

func main() {
	fmt.Print("enter int a=")
	fmt.Scanln(&a)
	fmt.Print("enter int b=")
	fmt.Scanln(&b)
	fmt.Println("sum =", add(a, b))
}
