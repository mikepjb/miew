package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	blackfriday "github.com/russross/blackfriday/v2"
)

func main() {
	fname := strings.Join(os.Args[1:], "")
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}

	output := blackfriday.Run(b)
	fmt.Printf(string(output))
}
