package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ascii2Decimal(ascii *string) {
	newASCII := strings.Split(*ascii, "")

	for _, val := range newASCII {
		_, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			fmt.Println(badInputError)
			return
		}
	}

	fmt.Printf("Decimal: ")
	for _, v := range newASCII {
		for _, v := range v {
			fmt.Printf("%d ", v)
		}
	}

}

func ascii2Hex(ascii *string) {
	newASCII := strings.Split(*ascii, "")

	for _, val := range newASCII {
		_, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			fmt.Println(badInputError)
			return
		}
	}

	fmt.Printf("Hexadecimal: ")
	for _, v := range newASCII {
		for _, v := range v {
			fmt.Printf("%x ", v)
		}
	}

}

func ascii2Binary(ascii *string) {
	newASCII := strings.Split(*ascii, "")

	for _, val := range newASCII {
		_, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			fmt.Println(badInputError)
			return

		}
	}

	fmt.Printf("Binary: ")
	for _, v := range newASCII {
		for _, v := range v {
			fmt.Printf("%b ", v)
		}
	}
}

func ascii2Octal(ascii *string) {
	newASCII := strings.Split(*ascii, "")

	for _, val := range newASCII {
		_, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			fmt.Println(badInputError)
			return
		}
	}

	fmt.Printf("Octal: ")
	for _, v := range newASCII {
		for _, v := range v {
			fmt.Printf("%o ", v)
		}
	}
}
