// chap1/manual-parse/main.go
package main

import (
	"fmt"
	"io"
	"os"
)

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
 
A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

// TODO – Insert definition of parseArgs() as earlier
// TODO – Insert definition of getName() as earlier
// TODO – Insert definition of greetUser() as earlier
// TODO – Insert definition of runCmd() as earlier

func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
