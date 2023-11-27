package tests

import (
	"testing"

	"github.com/perylemke/gowc/cmd"
)

func TestCountRunes(t *testing.T) {
	_, _, err := cmd.CountRunes("./testdata/test.txt")
	if err != nil {
		t.Errorf("Error on CountRunes function: %v", err)
	}
}
