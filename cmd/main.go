package main

import (
	"flag"

	"github.com/AnastasiiaDalakishvili/go-exercise/adapters"
	"github.com/AnastasiiaDalakishvili/go-exercise/domain"
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

	flag.Var(&inputFile, "input-file", "specify file name")
	flag.Parse()

	sum := domain.Add(inputFile)
	formattedSum := domain.FormatNumber(sum)
	adapters.Printer(formattedSum)
}
