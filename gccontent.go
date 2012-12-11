package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type DNAString struct {
	Label     string
	BasePairs string
}

type ParseState int

const (
	LABEL ParseState = iota
	PAIRS
)

func (dna DNAString) GCPercentage() float32 {
	numGC := 0
	for _, base := range dna.BasePairs {
		if base == 'G' || base == 'C' {
			numGC++
		}
	}
	//fmt.Println(dna, float32(numGC) / float32(len(dna.BasePairs))*100, len(dna.BasePairs), numGC)
	return float32(numGC) / float32(len(dna.BasePairs))
}

type DNAStringList []DNAString

func (dnalist DNAStringList) Len() int      { return len(dnalist) }
func (dnalist DNAStringList) Swap(i, j int) { dnalist[i], dnalist[j] = dnalist[j], dnalist[i] }

// for sorting
type ByGC struct{ DNAStringList }

func (byGC ByGC) Less(i, j int) bool {
	return byGC.DNAStringList[i].GCPercentage() > byGC.DNAStringList[j].GCPercentage()
}

func main() {
	r := bufio.NewReader(os.Stdin)

	results := make([]DNAString, 0)

	var tmp DNAString
	for {
		line, err := r.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		if len(line) > 0 && line[0] == '>' {
			if tmp.Label != "" {
				results = append(results, tmp)
				tmp.BasePairs = ""
			}
			tmp.Label = strings.TrimSpace(line[1:])
		} else {
			tmp.BasePairs = tmp.BasePairs + string(line)
		}
		if err != nil {
			results = append(results, tmp)
			break
		}
	}

	sort.Sort(ByGC{results})

	fmt.Printf("%s\n%2.8g%%\n", results[0].Label, results[0].GCPercentage()*100)
}
