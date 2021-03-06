package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var paths = []string{"./test.go.txt"}

var Expected = map[string][]string{
	"TestFail1":       {"tif", "tifif", "tfor", "tforfor", "trange", "tgo", "targ", "tclosure3", "tclosure2", "tclosure3", "<nil>", "tassign", "<nil>", "s", "sc", "<nil>"},
	"TestFail2":       {"tif", "tifif", "tfor", "tforfor", "trange", "tgo", "s", "sc"},
	"TestPass1":       {},
	"TestPass2":       {},
	"TestGCPartition": {"testKit"},
}

func TestFile(t *testing.T) {
	Cal = []string{"MustQuery", "MustExec", "mustExecute"}
	visitors, err := Calculate(paths)
	assert.Nil(t, err)
	assert.Len(t, visitors, 1)
	for _, visitor := range visitors {
		for funk, idents := range visitor.Occurrence {
			assert.Contains(t, Expected, funk.Name.String())

			expectedFunk := Expected[funk.Name.String()]
			assert.Len(t, idents, len(expectedFunk))
			for i := 0; i < len(idents); i++ {
				assert.Equal(t, expectedFunk[i], idents[i])
			}

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
