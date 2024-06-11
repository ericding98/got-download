package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/ericding98/got-download/src/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println(string(debug.Stack()))
		os.Exit(1)
	}
}
