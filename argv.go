// Argv shows its command line arguments.
package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func fmtArg(i int, s string) string {
	return fmt.Sprintf("%d (%d codepoints, %d bytes): %#v",
		i, utf8.RuneCountInString(s), len(s), s)
}

func main() {
	fmt.Println("argv length:", len(os.Args))

	for i, arg := range os.Args {
		fmt.Println(fmtArg(i, arg))
	}
}
