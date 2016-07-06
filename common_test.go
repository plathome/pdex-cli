package main

import (
	"testing"
	"github.com/plathome/pdex-cli/cmd"
)

func TestStringConcat(t *testing.T) {
	joined := StringConcat([]string {
		"foo",
		"bar",
	})
	if joined != "foobar" {
		t.Error("Strings concatenation error")
	}
}
