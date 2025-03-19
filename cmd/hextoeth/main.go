package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: hextoeth <hex>")
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

	res := big.NewFloat(0).SetInt(dec)
	res = res.Quo(res, big.NewFloat(1e18))

	fmt.Println(res)
}
