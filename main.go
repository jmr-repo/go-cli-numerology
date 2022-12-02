package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Hello Message
func HelloMessage() {
	// Prints Hello Message
	fmt.Println("Hello, Please enter your input to convert:")
	// Prints Cursor for User to type next to it
	fmt.Print("> ")
}

// Get User Input
func GetUserInput() []string {
	// Get User Input
	reader := bufio.NewReader(os.Stdin)

	// read string and store data into input variable
	input, err := reader.ReadString('\n')

	// check if err is nil
	if err != nil {
		// handle error
		fmt.Println("Error: ", err.Error())
	}

	// Remove spaces and other chars
	input = strings.TrimSuffix(input, "\n")

	// Convert string input to slice string
	s := strings.Split(input, "")

	// return []string
	return s
}

// Sum every number from slice
func SumSliceValues(s []string) int {
	var j int = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "A", "J", "S", "a", "j", "s":
			j = j + 1

		case "B", "K", "T", "b", "k", "t":
			j = j + 2

		case "C", "L", "U", "c", "l", "u":
			j = j + 3

		case "D", "M", "V", "d", "m", "v":
			j = j + 4

		case "E", "N", "W", "e", "n", "w":
			j = j + 5

		case "F", "O", "X", "f", "o", "x":
			j = j + 6

		case "G", "P", "Y", "g", "p", "y":
			j = j + 7

		case "H", "Q", "Z", "h", "q", "z":
			j = j + 8

		case "I", "R", "i", "r":
			j = j + 9

		}
	}

	return j
}

// Check if integer has more than 2 digits
func CheckNumDigits(i int) bool {
	if i > 9 {
		return true
	} else {
		return false
	}
}

// Convert int to string
func ConvertIntToString(i int) string {
	str := strconv.Itoa(i)
	return str
}

// Convert string to int
func ConvertStringToInt(s rune) int {
	i, err := strconv.Atoi(string(s))

	if err != nil {
		fmt.Println(err.Error())
	}

	return i
}

// Loop
func LoopSlice(i int) int {

	// set int
	var num int

	// Convert string argument into into
	str := ConvertIntToString(i)

	// loop through string ranges
	for _, digits := range str {

		// assign value strings as int
		value := ConvertStringToInt(digits)

		// add int to be returned
		num += value

	}

	// return int
	return num
}

// Logic
func Logic(i int) int {

	var r int

	if CheckNumDigits(i) {
		// if true ( > 9 )
		r = LoopSlice(i)

		if CheckNumDigits(r) {
			fmt.Println("First sum of your numbers is: ", r)
			r = LoopSlice(r)
		}

	} else {
		// if false ( < 9 )
		r = i
	}

	// return int
	return r
}

func main() {

	// Prints Hello Message
	HelloMessage()
	// Stores users input into a slice string var
	var input []string = GetUserInput()
	// Get int from given input
	num := SumSliceValues(input)

	if CheckNumDigits(num) {
		fmt.Println(":: Your full number is: ", num)
	}
	// Checks if int is greater than 9
	i := Logic(num)
	fmt.Println(":: Your Final Number: ", i)
}

// Table of values
//	1	2	3	4	5	6	7	8	9
//	A	B	C	D	E	F	G	H	I
//	J	K	L	M	N	O	P	Q	R
//	S	T	U	V	W	X	Y	Z
//	a	b	c	d	e	f	g	h	i
// 	j	k	l	m	n	o	p	q	r
//	s	t	u	v	w	x	y	z
