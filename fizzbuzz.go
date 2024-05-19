package main

import (
	"strconv"
)
func fizzBuzz(input int) string {

	mapFizz := make(map[bool]string, 2)
	mapFizz[true] = "Fizz"
	mapFizz[false] = ""

	mapBuzz := make(map[bool]string, 2)
	mapBuzz[true] = "Buzz"
	mapFizz[false] = ""

	if input % 3 == 0 || input % 5 == 0 {
		return mapFizz[input%3 == 0] + mapBuzz[input%5 == 0]
	}
	return strconv.Itoa(input)
}