package cmd

import (
	"fmt"
	"testing"
)

var paths = []string{"./test.go.txt"}

func TestFile(t *testing.T) {
	Cal = []string{"MustQuery", "MustExec", "mustExecute"}
	visitors, err := Calculate(paths)
	if err != nil {
		t.Fatal(err)
	}
	for _, visitor := range visitors {
		for funk, idents := range visitor.Occurrence {
			start := visitor.FileSet.Position(funk.Pos())
			end := visitor.FileSet.Position(funk.End())
			fmt.Printf("%s %v  total: %d from %d:%d to %d:%d\n", funk.Name, idents, len(idents), start.Line, start.Column, end.Line, end.Column)
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
