package main

import (
	"context"
	"fmt"
	start "world-pop/internal/init"
	"os"
)

var version string

func main() {
	ctx := context.Background()
	if err := start.Start(ctx, os.Stdout, os.Args, version); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
