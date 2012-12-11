package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for i, b := range in {
		if b == 'T' {
			in[i] = 'U'
		}
	}
	fmt.Println(string(in))
}
