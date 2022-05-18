package domain

import "github.com/AnastasiiaDalakishvili/go-exercise/adapters"

func Add(numbersPassedtoFunction ...interface{}) int {
	sum := 0
	if numbersPassedtoFunction != nil {
		//calculate sum
		for _, num := range numbersPassedtoFunction {
			if i, ok := num.(int); ok {
				sum += i
			}
		}
	} else {
		numbersFromFile := adapters.GetNumberFromFile()
		for _, num := range numbersFromFile {
			sum += num
		}
	}
	return sum
}