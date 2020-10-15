package utils

import (
	"os"

	"github.com/mattn/go-isatty"
)

var IsTerminal = func(f *os.File) bool {
	return isatty.IsTerminal(f.Fd()) || IsCygwinTerminal(f)
}

func IsCygwinTerminal(f *os.File) bool {
	return isatty.IsCygwinTerminal(f.Fd())
}
