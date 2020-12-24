package main

import (
	"fmt"

	"github.com/siuyin/present-basic_api_evolution/arith"
)

func main() {
	fmt.Println("Adding three numbers")
	fmt.Printf("2+3+4 = %v\n", arith.SumWithInterface(arith.ThreeNums{2, 3, 4}))

	fmt.Println("Adding four numbers")
	fourNum := arith.FourNumsEmb{
		ThreeNums: arith.ThreeNums{2, 3, 4},
		D:         5,
	}
	fmt.Printf("2+3+4+5 = %v\n", arith.SumWithInterface(fourNum))
}
