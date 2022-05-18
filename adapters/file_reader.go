package adapters

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AnastasiiaDalakishvili/go-exercise/helpers"
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

//Read the input file and return an array of numbers
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

	numbersFromFile = helpers.DataConverter(strings.Join(arrayOfInputs[:], ","))

	return numbersFromFile
}
