package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(checkNumberFromGenerator("12"))
	sumSelfNumber := 0
	for i := 1; i <= 5000; i++ {
		validateSelfNumber := checkSelfNumber(i)
		if validateSelfNumber {
			sumSelfNumber += i
		}
	}

	fmt.Println("Total: ")
	fmt.Println(sumSelfNumber)
}

func checkNumberFromGenerator(generator string) int {
	sum, _ := strconv.Atoi(generator)
	for _, val := range generator {
		numb, _ := strconv.Atoi(string(val))
		sum += numb
	}
	return sum
}

func checkSelfNumber(number int) bool {
	var res []int
	for i := number - 1; i > 0; i-- {
		strNumber := strconv.Itoa(i)
		result := checkNumberFromGenerator(strNumber)
		if result == number {
			res = append(res, result)
		}
	}

	if len(res) > 0 {
		return false
	}
	return true
}
