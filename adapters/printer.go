package adapters

import (
	"fmt"
	"io"
)

func Printer(w io.Writer, formattedSum string) {
	fmt.Fprintf(w, "Success %v\n", formattedSum)
}
