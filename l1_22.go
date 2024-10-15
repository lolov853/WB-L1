package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("2097152", 10)
	b.SetString("4194304", 10)

	sum := new(big.Int)
	diff := new(big.Int)
	product := new(big.Int)
	quotient := new(big.Int)

	sum.Add(a, b)
	diff.Sub(a, b)
	product.Mul(a, b)
	quotient.Quo(a, b)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
}
