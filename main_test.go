package main

import (
	"io"
	"os"
	"testing"

	"github.com/takumin/path-filter-action/internal/command"
)

func TestMainCommand(t *testing.T) {
	var capture int
	osExit = func(code int) { capture = code }

	ro, wo, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	re, we, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = wo
	os.Stderr = we

	os.Args = []string{"a"}

	main()

	wo.Close()
	we.Close()

	io.Copy(io.Discard, ro)
	io.Copy(io.Discard, re)

	expect := command.ExitOK
	actual := capture

	if expect != actual {
		t.Errorf(
			"Fail assert equal. Expect: %v Actual: %v",
			expect,
			actual,
		)
	}
}
