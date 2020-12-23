package main

import (
	"fmt"

	"github.com/siuyin/present-basic_api_evolution/arith"
)

func main() {
	fmt.Println("Adding three numbers")
	fmt.Printf("2+3+4 = %v\n", arith.SumGeneric(
		arith.GenericThing{Version: 1, X: arith.ThreeNums{2, 3, 4}}))

	fmt.Println("Adding four numbers")
	fmt.Printf("2+3+4+5 = %v\n", arith.SumGeneric(
		arith.GenericThing{Version: 2, X: arith.FourNums{2, 3, 4, 5}}))
}
