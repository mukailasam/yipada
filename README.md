### Yipada
Perfom conversions between ASCII, decimal, hexadecimal, octal, and binary values

### Install
go install github.com/mukailasam/yipada@latest

### Basic Usage

#### Decimal Conversion

Decimal to binary
`
yipada -db 10
`

Decimal to hexadecimal
`
yipada -dh 10
`

Decimal to ASCII
`
yipada -da 65
`

Decimal to octal
`
yipada -do 10
`

#### Binary Conversion

Binary to decimal
`
yipada -bd 1010
`

Binary to hexadecimal
`
yipada -bh 1010
`

Binary to ASCII
`
yipada -ba 1100001
`

Binary to octal
`
yipada -bo 1010
`
#### Hexadecimal Conversion

Hexadecimal to decimal
`
yipada -hd a
`

Hexadecimal to binary
`
yipada -hb a
`

Hexadecimal to ASCII
`
yipada -ha 41
`

Hexadecimal to octal
`
yipada -ho a
`

#### ASCII Conversion

ASCII to decimal
`
yipada -ad A
`

ASCII to binary
`
yipada -ab A
`

ASCII to hexadecimal
`
yipada -ah A
`

ASCII to octal
`
yipada -ao A
`

#### Octal Conversion

Octal to decimal
`
yipada -od 12
`

Octal to binary
`
yipada -ob 12
`

Octal to hexadecimal
`
yipada -oh 12
`

Octal to ASCII
`
yipada -oa 56
`

### Note
To perform conversion for multiple values, put the values in a double quote and seperate each value within the double quote with a space.
e.g
`
yidapa -ba "1101110 10110 1101001 110001"
`