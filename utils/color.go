package utils

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

var (
	// Output ANSI color if stdout is a tty
	Magenta = makeColorFunc("magenta")
	Cyan    = makeColorFunc("cyan")
	Red     = makeColorFunc("red")
	Yellow  = makeColorFunc("yellow")
	Blue    = makeColorFunc("blue")
	Green   = makeColorFunc("green")
	Gray    = makeColorFunc("black+b")
	Bold    = makeColorFunc("default+b")
)

func makeColorFunc(color string) func(string) string {
	cf := ansi.ColorFunc(color)
	return func(arg string) string {
		if isColorEnabled() {
			if color == "black+b" {
				return fmt.Sprintf("\x1b[%d;5;%dm%s\x1b[m", 38, 242, arg)
			}
			return cf(arg)
		}

		return arg
	}
}

func isColorEnabled() bool {
	return IsTerminal(os.Stdout)
}
