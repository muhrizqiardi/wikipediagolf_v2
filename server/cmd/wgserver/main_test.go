package main

import "testing"

func TestMain(t *testing.T) {
	var (
		exp = 1
		got = 1
	)

	if exp != got {
		t.Errorf("exp:%d; got:%d", exp, got)
	}
}
