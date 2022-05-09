package adapters

import (
	"fmt"
	"os"
)

//Read the input file and return a string
func GetNumberFromFile(filePath string) []int {
	numbersFromFile := []int{}

	//read data from the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//convert data into string
	input := string(data)

	numbersFromFile = DataConverter(input)

	return numbersFromFile
}
