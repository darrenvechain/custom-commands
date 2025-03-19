package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ethtohex <decimal>")
		os.Exit(1)
	}

	decimal := os.Args[1]
	eth, ok := big.NewInt(0).SetString(decimal, 10)
	if !ok {
		fmt.Println("Invalid decimal number")
		os.Exit(1)
	}

	eth = eth.Mul(eth, big.NewInt(1e18))

	fmt.Println("0x" + eth.Text(16))
}
