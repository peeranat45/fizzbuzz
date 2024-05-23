package main

import "strconv"



func fizzBuzz(input string) string {

	inputInt, err := strconv.Atoi(input)
	if err != nil {
		return input
	}

	fizzMap := make(map[bool]string, 2)
	fizzMap[true] = "Fizz"
	fizzMap[false] = ""

	buzzMap := make(map[bool]string, 2)
	buzzMap[true] = "Buzz"
	buzzMap[false] = ""

	inputMap := make(map[bool]string, 2)
	inputMap[true] = fizzMap[inputInt % 3 == 0] + buzzMap[inputInt % 5 == 0]
	inputMap[false] = input



	return inputMap[inputInt%3==0 || inputInt%5==0]
}