package subcmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestAtLeastTwo(t *testing.T) {
	// expected 	:= false
	// actual 		:= AtLeastTwo(true, true, false)

	// if actual != expected {
	// 	t.Errorf("Test failed, expected: '%b', got: '%b'", expected, actual)
	// }
}

// func TestRun_versionFlag(t *testing.T) {
// 	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
// 	cli 	:= &CLIoutStream: outStream, errStream: errStream

// 	fmt.Println(cli + "")
// 	// args 	:= strings.Split("pdex-cli -version"," ")

// 	// status := cli.Run(args)
// 	// if status != ExitCodeOK {
// 	// 	t.Errorf("ExitStatus=%d, want %d, status, ExitCodeOK")
// 	// }

// 	// expected := fmt.Sprintf("pdex-cli version %s", Version)

// 	// if !strings.Contains(errStream.String(), expected) {
// 	// 	t.Errorf("Output=%q, want %q", errStream.String(), expected)
// 	// }
// }

