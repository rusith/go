package learning

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io/ioutil"
)


func FindDuplicates() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	for line, c := range counts {
		if c > 0 {
			fmt.Printf("%s\t%d\n", line, c)
		}
	}
}

func FindDuplicatesFromFile() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, c := range counts {
		if c > 0 {
			fmt.Printf("%s\t%d\n", line, c)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if line == "exit" {
			break
		}
		counts[line]++
	}
}

func FindDuplicatesFromFileVersionTwo() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n" ) {
			counts[line]++
		}

	}

	for line, c := range counts {
		if c > 0 {
			fmt.Printf("%s\t%d\n", line, c)
		}
	}
}