package tests

import (
	"testing"

	"github.com/perylemke/gowc/cmd"
)

func TestCountWords(t *testing.T) {
	_, _, err := cmd.CountWords("./testdata/test.txt")
	if err != nil {
		t.Errorf("Error on CountWords function: %v", err)
	}
}
