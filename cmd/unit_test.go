package cmd

import (
	"testing"
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

