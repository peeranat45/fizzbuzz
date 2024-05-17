package main

import (
	"fmt"
	"log"
	"strconv"
)

func fizzBuzz(input string) string {

	inputInt , err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if  inputInt%3 == 0 {
		return "Fizz"
	}else if inputInt%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf(input)
}