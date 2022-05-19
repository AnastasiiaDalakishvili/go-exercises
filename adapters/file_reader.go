package adapters

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AnastasiiaDalakishvili/go-exercise/helpers"
)

//files
type arrayOfFlags []string

func (f *arrayOfFlags) String() string {
	return ""
}

func (f *arrayOfFlags) Set(flag string) error {
	*f = append(*f, flag)
	return nil
}

var arrayOfFileNamesFromCLI arrayOfFlags

//numbers
type arrayOfNumbers []string

func (n *arrayOfNumbers) String() string {
	return ""
}

func (n *arrayOfNumbers) Set(flag string) error {
	*n = append(*n, flag)
	return nil
}

var arrayOfNumbersFromCLI arrayOfNumbers

func GetDataFromCLI() []int {
	var numbers []int

	//get numbers passed to cli
	flag.Var(&arrayOfNumbersFromCLI, "input-numbers", "pass numbers")

	//get files passed to cli
	flag.Var(&arrayOfFileNamesFromCLI, "input-file", "pass file name")

	flag.Parse()

	if len(arrayOfNumbersFromCLI) != 0 {
		numbers = GetNumbersFromCLI()
	} else if len(arrayOfFileNamesFromCLI) != 0 {
		numbers = GetNumbersFromFile()
	}
	return numbers
}

func GetNumbersFromCLI() []int {
	var numbers []int
	var arrayOfNumbers []string

	for _, str := range arrayOfNumbersFromCLI {
		arrayOfNumbers = strings.Split(str, ",")
	}

	for _, i := range arrayOfNumbers {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, j)
	}

	return numbers
}

func GetNumbersFromFile() []int {
	var numbers []int
	var arrayOfInputs []string

	for _, file := range arrayOfFileNamesFromCLI {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		input := string(data)
		arrayOfInputs = append(arrayOfInputs, input)
	}

	numbers = helpers.DataConverter(strings.Join(arrayOfInputs[:], ","))

	return numbers
}
