package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var complements = map[byte]byte{
	'A': 'T',
	'T': 'A',
	'G': 'C',
	'C': 'G',
}

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for i, b := range in {
		in[i] = complements[b]
	}
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	fmt.Println(string(in))
}
