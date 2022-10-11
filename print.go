package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

var cyan = color.CyanString
var blue = color.BlueString
var yellow = color.YellowString
var green = color.GreenString
var white = color.WhiteString
var red = color.RedString

func write(format string, args ...interface{}) {
	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf(format, args...))
}

func writeEmptyLine() {
	write("")
}
