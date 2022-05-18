package adapters

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

//Read the input file and return a string
func GetNumberFromFile() []int {

	flag.Var(&inputFile, "input-file", "specify file name")
	flag.Parse()

	numbersFromFile := []int{}
	var arrayOfInputs []string

	for _, file := range inputFile {
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
