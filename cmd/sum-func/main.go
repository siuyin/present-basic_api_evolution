package main

import "fmt"

func main() {
	fmt.Println("Adding two numbers")
	fmt.Printf("2+3 = %v\n", sum(2, 3))
}

func sum(a, b int) int {
	return a + b
}
