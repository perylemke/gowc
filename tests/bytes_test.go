package tests

import (
	"testing"

	"github.com/perylemke/gowc/cmd"
)

func TestCountBytes(t *testing.T) {
	_, _, err := cmd.CountBytes("./testdata/test.txt")
	if err != nil {
		t.Errorf("Error on CountBytes function: %v", err)
	}
}
