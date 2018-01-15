package ioutils

import (
	"fmt"
	"os"
)

func Errorf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func Error(msg interface{}) {
	Errorf("%s\n", msg)
}