package main

import (
	"flag"
	"os"
)

func main() {

	// Decimal
	d2b := flag.String("db", "", "Decimal to binay")
	d2h := flag.String("dh", "", "Decimal to hexadecimal")
	d2a := flag.String("da", "", "Decimal to ASCII")
	d2o := flag.String("do", "", "Decimal to octal")

	// Binary
	b2d := flag.String("bd", "", "Binary to decimal")
	b2h := flag.String("bh", "", "Binary to hexadecimal")
	b2a := flag.String("ba", "", "Binary to ASCII")
	b2o := flag.String("bo", "", "Binary to octal")

	// Hexadecimal
	h2d := flag.String("hd", "", "Hexadecimal to decimal")
	h2b := flag.String("hb", "", "Hexadecimal to binary")
	h2a := flag.String("ha", "", "Hexadecimal to ASCII")
	h2o := flag.String("ho", "", "Hexadecimal to octal")

	// ASCII
	a2d := flag.String("ad", "", "ASCII to decimal")
	a2b := flag.String("ab", "", "ASCII to binary")
	a2h := flag.String("ah", "", "ASCII to hexadecimal")
	a2o := flag.String("ao", "", "ASCII to octal")

	// Octal
	o2d := flag.String("od", "", "Octal to decimal")
	o2b := flag.String("ob", "", "Octal to binary")
	o2h := flag.String("oh", "", "Octal to hexadecimal")
	o2a := flag.String("oa", "", "Octal to ASCII")

	flag.Parse()

	args := os.Args[:]

	if len(args) <= 1 {
		flag.Usage()
		os.Exit(1)
	}

	switch args[1] {
	// Decimal
	case "-db":
		decimal2Binary(d2b)
	case "-dh":
		decimal2Hex(d2h)
	case "-da":
		decimal2ASCII(d2a)
	case "-do":
		decimal2Octal(d2o)

	// Binary
	case "-bd":
		binary2Decimal(b2d)
	case "-bh":
		binary2Hex(b2h)
	case "-ba":
		binary2ASCII(b2a)
	case "-bo":
		binary2Octal(b2o)

	// Hexadecimal
	case "-hd":
		hex2Decimal(h2d)
	case "-hb":
		hex2Binary(h2b)
	case "-ha":
		hex2ASCII(h2a)
	case "-ho":
		hex2Octal(h2o)

	// ASCII
	case "-ad":
		ascii2Decimal(a2d)
	case "-ab":
		ascii2Binary(a2b)
	case "-ah":
		ascii2Hex(a2h)
	case "-ao":
		ascii2Octal(a2o)

	// Octal
	case "-od":
		octal2Decimal(o2d)
	case "-ob":
		octal2Binary(o2b)
	case "-oh":
		octal2Hex(o2h)
	case "-oa":
		octal2ascii(o2a)

	//default
	default:
		flag.Usage()
	}
}
