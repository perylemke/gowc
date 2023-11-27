package tests

import (
	"testing"

	"github.com/perylemke/gowc/cmd"
)

func TestCountLines(t *testing.T) {
	_, _, err := cmd.CountLines("./testdata/test.txt")
	if err != nil {
		t.Errorf("Error on CountLines function: %v", err)
	}
}
