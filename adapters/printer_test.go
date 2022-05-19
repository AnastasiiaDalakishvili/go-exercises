package adapters

import (
	"bytes"
	"testing"
)

func TestPrinter(t *testing.T) {
	t.Skip()
	expect := "Success 123"

	t.Run("testing", func(t *testing.T) {
		var output bytes.Buffer

		Printer(&output, "123")
		if expect != output.String() {
			t.Errorf("got %s expected %s", output.String(), expect)
		}
	})
}
