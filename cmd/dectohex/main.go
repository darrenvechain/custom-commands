package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: dectohex <decimal>")
		os.Exit(1)
	}

	decimal := os.Args[1]
	val := big.NewInt(0)
	val, ok := val.SetString(decimal, 10)
	if !ok {
		fmt.Println("Invalid decimal number")
		os.Exit(1)
	}

	fmt.Println("0x" + val.Text(16))
}
