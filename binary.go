package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binary2Hex(binary *string) {
	newBinary := strings.Split(*binary, " ")

	output := make([]int64, 0)
	for _, nb := range newBinary {
		nv, err := strconv.ParseInt(nb, 2, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nv)
	}

	fmt.Printf("Hexadecimal: ")
	for _, val := range output {
		fmt.Printf("%x ", val)
	}
}

func binary2Decimal(binary *string) {
	newBinary := strings.Split(*binary, " ")

	output := make([]int64, 0)
	for _, nb := range newBinary {
		nv, err := strconv.ParseInt(nb, 2, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nv)
	}

	fmt.Printf("Decimal: ")
	for _, val := range output {
		fmt.Printf("%d ", val)
	}
}

func binary2Octal(binary *string) {
	newBinary := strings.Split(*binary, " ")

	output := make([]int64, 0)
	for _, nb := range newBinary {
		nv, err := strconv.ParseInt(nb, 2, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nv)
	}

	fmt.Printf("Octal: ")
	for _, val := range output {
		fmt.Printf("%o ", val)
	}

}

func binary2ASCII(binary *string) {
	newBinary := strings.Split(*binary, " ")

	output := make([]int64, 0)
	for _, nb := range newBinary {
		nv, err := strconv.ParseInt(nb, 2, 64)
		if err != nil {
			fmt.Println(badInputError)
			return
		}

		output = append(output, nv)
	}

	fmt.Printf("ASCII: ")
	for _, val := range output {
		fmt.Printf("%c ", val)
	}

}
