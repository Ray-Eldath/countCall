package cmd

import (
	"fmt"
	"testing"
)

var paths = []string{"/Users/ray/Library/Application Support/JetBrains/GoLand2021.2/scratches/test.go"}

func TestFile(t *testing.T) {
	visitors, err := Calculate(paths)
	if err != nil {
		t.Fatal(err)
	}
	for _, visitor := range visitors {
		for funk, idents := range visitor.Occurrence {
			fmt.Printf("%s %v  total: %d\n", funk.Name, idents, len(idents))
		}
	}
}

func TestPrint(t *testing.T) {
	err := Print(paths)
	if err != nil {
		t.Fatal(err)
		return
	}
}
