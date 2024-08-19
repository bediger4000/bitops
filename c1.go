package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	n, err := strconv.ParseUint(os.Args[1], 2, 8)
	if err != nil {
		log.Fatal(err)
	}

	x := uint8(n)
	even := x & 0b01010101
	odd := x & 0b10101010
	swapped := (even << 1) | (odd >> 1)

	fmt.Printf("original: %08b\n", x)
	fmt.Printf("swapped:  %08b\n", swapped)

	swapped = ((x & 0b01010101) << 1) | ((x & 0b10101010) >> 1)

	fmt.Printf("swapped:  %08b\n", swapped)
}
