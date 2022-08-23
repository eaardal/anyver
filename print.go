package main

import (
	"fmt"
	"os"
)

func write(format string, args ...interface{}) {
	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf(format, args...))
}

func writeEmptyLine() {
	write("")
}
