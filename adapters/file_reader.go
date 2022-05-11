package adapters

import (
	"fmt"
	"os"
	"strings"
)

//Read the input file and return a string
func GetNumberFromFile(filePath []string) []int {
	numbersFromFile := []int{}
	var arrayOfInputs []string

	for _, file := range filePath {
		//read data from the file

		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//convert data into string
		input := string(data)

		arrayOfInputs = append(arrayOfInputs, input)
	}

	numbersFromFile = DataConverter(strings.Join(arrayOfInputs[:], ","))

	return numbersFromFile
}
