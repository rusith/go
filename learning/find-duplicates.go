package learning

import (
	"fmt"
	"os"
	"bufio"
)


func FindDuplicates() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "exit" {
			break
		}

		counts[line]++
	}

	for line, c := range counts {
		if c > 0 {
			fmt.Printf("%s\t%d\n", line, c)
		}
	}
}