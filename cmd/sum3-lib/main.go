package main

import (
	"fmt"

	"github.com/siuyin/present-basic_api_evolution/arith"
)

func main() {
	fmt.Println("Adding two numbers")
	fmt.Printf("2+3 = %v\n", arith.Sum(2, 3))

	fmt.Println("Adding three numbers")
	fmt.Printf("2+3+4 = %v\n", arith.Sum3(2, 3, 4))
}
