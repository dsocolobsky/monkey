package main

import (
	"fmt"
	"os"

	"github.com/dysoco/monkey/repl"
)

func main() {
	fmt.Printf("Welcome to monkey")
	repl.Start(os.Stdin, os.Stdout)
}
