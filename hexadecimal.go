package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hex2Binary(hex *string) {
	newHex := strings.Split(*hex, " ")

	output := make([]int64, 0)
	for _, nv := range newHex {
		nd, err := strconv.ParseInt(nv, 16, 64)
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

func hex2Octal(hex *string) {
	newHex := strings.Split(*hex, " ")

	output := make([]int64, 0)
	for _, nv := range newHex {
		nd, err := strconv.ParseInt(nv, 16, 64)
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

func hex2ASCII(hex *string) {
	newHex := strings.Split(*hex, " ")

	output := make([]int64, 0)
	for _, nv := range newHex {
		nd, err := strconv.ParseInt(nv, 16, 64)
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

func hex2Decimal(hex *string) {
	newHex := strings.Split(*hex, " ")

	output := make([]int64, 0)
	for _, nv := range newHex {
		nd, err := strconv.ParseInt(nv, 16, 64)
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
