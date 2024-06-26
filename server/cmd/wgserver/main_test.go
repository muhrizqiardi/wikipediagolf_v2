package main

import (
	"context"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	var (
		ctx    = context.TODO()
		args   = []string{"wgserver"}
		getenv = func(string) string { return "" }
		stdin  = os.Stdin
		stdout = os.Stdout
		stderr = os.Stdout
	)

	if err := run(ctx, args, getenv, stdin, stdout, stderr); err != nil {
		t.Error("exp nil; got err:", err)
	}
}
