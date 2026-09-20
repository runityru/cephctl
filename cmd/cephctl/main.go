package main

import (
	"context"
	"os"
)

func main() {
	// ExecuteContext runs the cobra command tree.
	//
	// SilenceUsage=true: usage is not printed when a RunE handler returns an
	// error. SilenceErrors=false (default): cobra prints the error text to
	// stderr (ErrOrStderr), after which we exit with code 1.
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
