package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	blackfriday "github.com/russross/blackfriday/v2"
)

// handler is not a good name.
func handler(w http.ResponseWriter, r *http.Request) {
	fname := strings.Join(os.Args[1:], "")
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		fmt.Printf("error reading file: %v", err)
	}

	content := blackfriday.Run(b)

	w.Write(content)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
