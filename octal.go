package main

import (
	"fmt"
	"strconv"
	"strings"
)

func octal2Decimal(octal *string) {
	newOctal := strings.Split(*octal, " ")

	output := make([]int64, 0)
	for _, no := range newOctal {
		nd, err := strconv.ParseInt(no, 8, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nd)

	}

	fmt.Printf("Decimal: ")
	for _, val := range output {
		fmt.Printf("%d ", val)
	}
}

func octal2Hex(octal *string) {
	newOctal := strings.Split(*octal, " ")

	output := make([]int64, 0)
	for _, no := range newOctal {
		nd, err := strconv.ParseInt(no, 8, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nd)
	}

	fmt.Printf("Hexadecimal: ")
	for _, val := range output {
		fmt.Printf("%x ", val)
	}
}

func octal2Binary(octal *string) {
	newOctal := strings.Split(*octal, " ")

	output := make([]int64, 0)
	for _, no := range newOctal {
		nd, err := strconv.ParseInt(no, 8, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nd)
	}

	fmt.Printf("Binary: ")
	for _, val := range output {
		fmt.Printf("%b ", val)
	}
}

func octal2ascii(octal *string) {
	newOctal := strings.Split(*octal, " ")

	output := make([]int64, 0)
	for _, no := range newOctal {
		nd, err := strconv.ParseInt(no, 8, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nd)

	}

	fmt.Printf("ASCII: ")
	for _, val := range output {
		fmt.Printf("%c ", val)
	}
}
