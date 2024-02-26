package main

import (
	"cantolang/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
