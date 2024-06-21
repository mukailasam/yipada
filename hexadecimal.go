package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

func hex2Binary(h *string) {
	output, err := strconv.ParseInt(*h, 16, 64)
	if err != nil {
		log.Fatalln(badInputError)
	}
	fmt.Printf("Binary: %b", output)
}

func hex2Octal(h *string) {
	output, err := strconv.ParseInt(*h, 16, 64)
	if err != nil {
		log.Fatalln(badInputError)
	}
	fmt.Printf("Octal: %o", output)
}

func hex2ASCII(h *string) {
	output, err := hex.DecodeString(*h)
	if err != nil {
		log.Fatalln(badInputError)
	}
	fmt.Printf("ASCII: %s", string(output))
}

func hex2Decimal(h *string) {
	output, err := strconv.ParseInt(*h, 16, 64)
	if err != nil {
		log.Fatalln(badInputError)
	}
	fmt.Printf("Decimal: %d", output)
}
