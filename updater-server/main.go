package main

import (
	"context"
	"fmt"
	"os"
	start "updater-server/internal/init"
)

func main() {
	ctx := context.Background()
	if err := start.Start(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
