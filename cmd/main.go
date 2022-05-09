package main

import (
	"flag"
	"fmt"

	"github.com/AnastasiiaDalakishvili/go-exercise/adapters"
)

func main() {
	// Declare a string flag called file path with a default value ("input.txt") and a help message
	fileName := flag.String("input-file", "", "specify file name")
	flag.Parse()

	path := *fileName

	result := Add(path)
	fmt.Println(formatNumber(result))
}

//Add calculates the sum of the numbers passed to it. If no data passed to the function then it reads it from the file
func Add(filePath string, numbersPassedtoFunction ...interface{}) int {
	sum := 0
	if numbersPassedtoFunction != nil {
		//calculate sum
		for _, num := range numbersPassedtoFunction {
			if i, ok := num.(int); ok {
				sum += i
			}
		}
	} else {
		numbersFromFile := adapters.GetNumberFromFile(filePath)
		for _, num := range numbersFromFile {
			sum += num
		}
	}
	return sum
}
