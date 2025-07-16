package main

import (
	"context"
	"fmt"
	"github.com/eric-schulze/self-updater/internal/init"
	"os"
)

func main() {
	ctx := context.Background()
	if err := init.Start(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}