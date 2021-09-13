package main

import (
	"flag"
	"fmt"
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

func main() {
	flag.Var(&filesFlag, "file", "A single file that you want to calculate. Set this multiple times if you have many files.")
	flag.Parse()
	if len(filesFlag) <= 0 {
		fmt.Println("Error: no file specified, check out --help for usage.")
		os.Exit(255)
	}

	visitors, err := cmd.Calculate(filesFlag)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	for path, visitor := range visitors {
		for funk, idents := range visitor.Occurrence {
			fmt.Printf("%s,%s,%d\n", path, funk.Name.String(), len(idents))
		}
	}
}
