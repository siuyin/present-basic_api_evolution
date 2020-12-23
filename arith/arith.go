// 10 OMIT
// Package arith provides arithmetic functions.
package arith

import "log"

// Sum returns the sum of integers a and b.
func Sum(a, b int) int {
	return a + b
}

// 20 OMIT

// Sum3 returns the sum of integers a, b and c.
func Sum3(a, b, c int) int {
	return a + b + c
}

// 30 OMIT

// ThreeNums contains 3 integers.
type ThreeNums struct {
	A, B, C int
}

// FourNums contains 4 integers.
type FourNums struct {
	A, B, C, D int
}

// SumGP takes an empty interface x, sums its
// contents and returns an int.
func SumGP(x interface{}) int {
	switch v := x.(type) {
	case ThreeNums:
		return v.A + v.B + v.C

	case FourNums:
		return v.A + v.B + v.C + v.D
	default:
		log.Fatal("help! I don't know how to sum this!")
	}
	return -999999 // BAD: return an error in Go instead
}

// 40 OMIT
// GenericThing is a versioned type,
// X can be "any one thing"/empty interface.
type GenericThing struct {
	Version int
	X       interface{}
}

func SumGeneric(y GenericThing) int {
	switch y.Version {
	case 1:
		v := y.X.(ThreeNums) // assert that v must be a ThreeNums
		return v.A + v.B + v.C
	case 2:
		v := y.X.(FourNums) // assert that v must be a FourNums
		return v.A + v.B + v.C + v.D
	default:
		log.Fatal("help! I don't know how to sum this!")
	}
	return -999999 // BAD: return an error in Go instead
}

// 50 OMIT
