package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: hextodec <hex>")
		os.Exit(1)
	}

	hex := os.Args[1]
	if strings.HasPrefix(hex, "0x") {
		hex = hex[2:]
	}

	dec, ok := big.NewInt(0).SetString(hex, 16)
	if !ok {
		fmt.Println("Invalid hex number")
		os.Exit(1)
	}

	fmt.Println(dec)
}
