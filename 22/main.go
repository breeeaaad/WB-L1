package main

import (
	"fmt"
	"math/big"
)

func main() {
	f := big.NewInt(20000000000)
	s := big.NewInt(20000000000)

	add := new(big.Int)
	fmt.Println(add.Add(f, s))

	sub := new(big.Int)
	fmt.Println(sub.Sub(f, s))

	quo := new(big.Int)
	fmt.Println(quo.Quo(f, s))
}
