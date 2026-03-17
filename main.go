package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	prefill := strings.Join(os.Args[1:], " ")

	cmd, err := run(prefill)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if cmd != "" {
		fmt.Print(cmd)
	}
}
