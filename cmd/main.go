package main

import (
	"flag"
	"fmt"

	"github.com/AnastasiiaDalakishvili/go-exercise/adapters"
)

type arrayOfFlags []string

func (f *arrayOfFlags) String() string {
	return ""
}

func (f *arrayOfFlags) Set(flag string) error {
	*f = append(*f, flag)
	return nil
}

var inputFile arrayOfFlags

func main() {
	// Declare a string flag called file path with a default value ("input.txt") and a help message
	//fileName := flag.String("input-file", "", "specify file name")
	flag.Var(&inputFile, "input-file", "specify file name")
	flag.Parse()

	//convertedInput := adapters.DataConverter(inputFile)

	//path := *fileName

	result := Add(inputFile)
	fmt.Println(result)
	//fmt.Println(domain.FormatNumber(result))
}

//Add calculates the sum of the numbers passed to it. If no data passed to the function then it reads it from the file
func Add(filePath []string, numbersPassedtoFunction ...interface{}) int {
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
