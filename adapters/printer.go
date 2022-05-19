package adapters

import (
	"fmt"
	"io"
)

//func Printer(formattedSum string) {
//	fmt.Printf("Success %v\n", formattedSum)
//}

func Printer(w io.Writer, formattedSum string) {
	fmt.Fprintf(w, "Success %v\n", formattedSum)
}

//func SayHello(w io.Writer, name string) {
//	fmt.Fprintf(w, "Hi, my name is %s\n", name)
//}
