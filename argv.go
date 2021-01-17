// Argv prints its command line arguments.
package main

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

func fmtArg(i int, s string) string {
	return fmt.Sprintf("%d (%d codepoints, %d bytes): %#v",
		i, utf8.RuneCountInString(s), len(s), s)
}

func writeArgv(argv []string, w io.Writer) {
	fmt.Fprintln(w, "argv length:", len(argv))

	for i, arg := range argv {
		fmt.Fprintln(w, fmtArg(i, arg))
	}
}

func main() {
	writeArgv(os.Args, os.Stdout)
}
