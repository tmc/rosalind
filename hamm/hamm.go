package main

import (
	"bufio"
	"fmt"
	"os"
)

type pairs []byte

func HammingDistance(a, b pairs) (result int, err error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("lengths mismatched: %d and %d", len(a), len(b))
	}
	for i, _ := range a {
		if a[i] != b[i] {
			result++
		}
	}
	return result, nil
}

func main() {
	r := bufio.NewReader(os.Stdin)

	line1, _ := r.ReadString('\n')
	line2, _ := r.ReadString('\n')

	hd, err := HammingDistance(pairs(line1), pairs(line2))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hd)
}
