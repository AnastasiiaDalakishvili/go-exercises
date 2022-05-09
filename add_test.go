package main

import (
	"testing"
)

const TestFilePassTxt = "data/input1.txt"
const TestFilePassCsv = "data/input2.csv"

func TestAdd(t *testing.T) {
	t.Run("Should take any number of arguments and print out the sum. Ignore any non-integers", func(t *testing.T) {
		sum := Add(TestFilePassTxt, 5, 4, 2, -10, 4.2, "l", 32, 14)

		got := formatNumber(sum)
		want := "47"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Should return number without commas", func(t *testing.T) {
		sum := Add(TestFilePassTxt, 9999)

		got := formatNumber(sum)
		want := "9999"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Add should take any number of arguments and print out the sum. Ignore any non-integers", func(t *testing.T) {
		sum := Add(TestFilePassTxt, 5, 4, 2, -10, 4.2, "l", 9952, 8567292)

		got := formatNumber(sum)
		want := "8,577,245"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("If no arguments passed to the function it should read data from the txt file", func(t *testing.T) {
		sum := Add(TestFilePassTxt)

		got := formatNumber(sum)
		want := "867,684"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("If no arguments passed to the function it should read data from the csv file", func(t *testing.T) {
		sum := Add(TestFilePassCsv)

		got := formatNumber(sum)
		want := "867,685"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}