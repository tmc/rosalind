package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	counts := map[byte]int{'A': 0, 'T': 0, 'G': 0, 'C': 0}
	i, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, b := range i {
		counts[b]++
	}
	fmt.Println(counts['A'], counts['C'], counts['G'], counts['T'])
}
