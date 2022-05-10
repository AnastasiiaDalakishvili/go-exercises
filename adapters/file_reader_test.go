package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestFilePassTxt = "/Users/anastasiia.dalakishvili/github/saltpay/go-exercises/go-exercises/data/input1.txt"
const TestFilePassCsv = "/Users/anastasiia.dalakishvili/github/saltpay/go-exercises/go-exercises/data/input2.csv"

func TestGetNumberFromFile(t *testing.T) {
	t.Run("Converter should return an array of integers", func(t *testing.T) {
		files := []string{TestFilePassTxt, TestFilePassCsv}

		desiredRes := []int{}

		result := GetNumberFromFile(files)

		assert.IsType(t, desiredRes, result)

	})
}
