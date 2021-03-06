package main

import (
	"flag"
	"fmt"
	"go/token"
	"os"
	"strings"

	"calCall/cmd"
)

type arrayFlags []string

func (i *arrayFlags) String() string { return strings.Join(*i, " ") }
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var filesFlag arrayFlags
var calFlag arrayFlags

func main() {
	flag.Var(&calFlag, "cal", "The name of func that you want to count. Set this multiple times if you have many keywords.")
	flag.Var(&filesFlag, "file", "A single file that you want to analyze. Set this multiple times if you have many files.")
	flag.Parse()
	if len(filesFlag) <= 0 {
		fmt.Println("Error: no file specified, check out --help for usage.")
		os.Exit(255)
	}
	if len(calFlag) <= 0 {
		fmt.Println("Error: no cal specified, check out --help for usage.")
		os.Exit(255)
	}

	cmd.Cal = calFlag
	visitors, err := cmd.Calculate(filesFlag)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	for path, v := range visitors {
		for funk, idents := range v.Occurrence {
			fmt.Printf("%s,%s,%s,%s,%d\n", path, Position(v, funk.Pos()), Position(v, funk.End()), funk.Name.String(), len(idents))
		}
	}
}

func Position(v *cmd.Visitor, pos token.Pos) string {
	t := v.FileSet.Position(pos)
	return fmt.Sprintf("%d:%d", t.Line, t.Column)
}
