package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decimal2Binary(decimal *string) {
	newDecimal := strings.Split(*decimal, " ")

	output := make([]int64, 0)
	for _, nv := range newDecimal {
		nd, err := strconv.ParseInt(nv, 10, 64)
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

func decimal2ASCII(decimal *string) {
	newDecimal := strings.Split(*decimal, " ")

	output := make([]int64, 0)
	for _, nv := range newDecimal {
		nd, err := strconv.ParseInt(nv, 10, 64)
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

func decimal2Hex(decimal *string) {
	newDecimal := strings.Split(*decimal, " ")

	output := make([]int64, 0)
	for _, nv := range newDecimal {
		nd, err := strconv.ParseInt(nv, 10, 64)
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

func decimal2Octal(decimal *string) {
	newDecimal := strings.Split(*decimal, " ")

	output := make([]int64, 0)
	for _, nv := range newDecimal {
		nd, err := strconv.ParseInt(nv, 10, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nd)

	}

	fmt.Printf("Octal: ")
	for _, val := range output {
		fmt.Printf("%o ", val)
	}
}
