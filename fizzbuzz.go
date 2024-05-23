package main

import "strconv"



func fizzBuzz(input string) string {

	inputInt, err := strconv.Atoi(input)
	if err != nil {
		return input
	}
	fizzMap := make(map[bool]string, 2)

	fizzMap[true] = "Fizz"
	fizzMap[false] = input
	return fizzMap[inputInt % 3 == 0]
}