package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	x, err := strconv.ParseUint(os.Args[1], 0x10, 64)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseUint(os.Args[2], 0x10, 64)
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.ParseUint(os.Args[3], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if !(b == 0 || b == 1) {
		log.Fatalf("b has to have 0 or 1 value, not %d\n", b)
	}
	fmt.Printf("%x\t%x\t%d\n", x, y, b)

	var zork uint64
	zork = 0xFFFFFFFFFFFFFFFF

	result := y & (zork + b)
	result += x & uint64(0-b)

	fmt.Printf("result: %x\n", result)
}
