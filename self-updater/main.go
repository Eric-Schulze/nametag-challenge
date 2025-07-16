package main

import (
	"context"
	"fmt"
	"github.com/eric-schulze/nametag-challeng/self-updater/internal/bootstrap"
	"os"
)

func main() {
	ctx := context.Background()
	if err := bootstrap.Start(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}